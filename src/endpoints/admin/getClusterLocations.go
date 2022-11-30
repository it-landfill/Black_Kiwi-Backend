package black_kiwi_admin

import (
	"net/http"
	"time"
	"os"
	"strconv"


	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/dbHandler/admin"
	"ITLandfill/Black-Kiwi/structs/data_structs"
)

// getAlbums responds with the list of all albums as JSON.
func GetClusterLocations(c *gin.Context) {
	from := c.DefaultQuery("from", "")
	to := c.DefaultQuery("to", "")
	
	nClusterStr := c.DefaultQuery("nCluster", "")
	if nClusterStr == "" {
		log.Warn("Missing nCluster parameter in GetClusterLocations")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing nCluster parameter"})
		return
	}
	nCluster, err := strconv.Atoi(nClusterStr)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error while parsing nCluster in GetClusterLocations")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if nCluster <= 0 {
		log.WithFields(log.Fields{"error": err}).Error("nCluster must be > 0 in GetClusterLocations")
		c.JSON(http.StatusBadRequest, gin.H{"error": "nCluster must be > 0"})
		return
	}

	var maxRadius float32 = 0
	maxRadiusStr := c.DefaultQuery("maxRadius", "")
	if maxRadiusStr != "" {
		maxRadiusTmp, err := strconv.ParseFloat(maxRadiusStr, 32)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("Error while parsing maxRadius in GetClusterLocations")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if maxRadiusTmp <= 0 {
			log.WithFields(log.Fields{"error": err}).Error("maxRadius must be > 0 in GetClusterLocations")
			c.JSON(http.StatusBadRequest, gin.H{"error": "maxRadius must be > 0"})
			return
		}	
		maxRadius = float32(maxRadiusTmp)
	}

	var success bool
	var requestList *[]black_kiwi_data_structs.ClusterInfo 

	if os.Getenv("Black_Kiwi_ENV") == "dev-nodb" { 
		log.WithFields(log.Fields{"endpoint": "GetClusterLocations"}).Info("Endpoint called in dev-nodb mode")
		c.IndentedJSON(http.StatusOK, black_kiwi_data_structs.MockClusterInfo)
		return
	}

	if from != "" && to != "" {
		fromTime, err := time.Parse("2006-01-02", from)
		if err != nil {
			log.WithFields(log.Fields{"from": from}).Warn("Unable to parse from time in GetClusterLocations")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse from time in GetClusterLocations"})
			return
		}

		toTime, err := time.Parse("2006-01-02", to)
		if err != nil {
			log.WithFields(log.Fields{"to": to}).Warn("Unable to parse to time in GetClusterLocations")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse to time in GetClusterLocations"})
			return
		}

		log.WithFields(log.Fields{"endpoint": "GetClusterLocations", "from": fromTime, "to": toTime, "maxRadius": maxRadius, "nCluster": nCluster}).Info("GetClusterLocations endpoint called")
		success, requestList = black_kiwi_admin_queries.GetClusterLocation(&fromTime, &toTime, nCluster, maxRadius)
	} else {
		log.WithFields(log.Fields{"endpoint": "GetClusterLocations"}).Info("GetClusterLocations endpoint called")
		success, requestList = black_kiwi_admin_queries.GetClusterLocation(nil, nil, nCluster, maxRadius)
	}

	if !success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting cluster locations"})
		return
	}

	if requestList == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No requests found"})
		return
	}

	c.IndentedJSON(http.StatusOK, *requestList)
}
