package black_kiwi_admin

import (
	"net/http"
	"time"
	"os"


	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/dbHandler/admin"
	"ITLandfill/Black-Kiwi/structs/data_structs"
)

// getAlbums responds with the list of all albums as JSON.
func GetRequestLocations(c *gin.Context) {
	from := c.DefaultQuery("from", "")
	to := c.DefaultQuery("to", "")

	var success bool
	var requestList *[]black_kiwi_data_structs.RequestInfo 

	if os.Getenv("Black_Kiwi_ENV") == "dev-nodb" { 
		log.WithFields(log.Fields{"endpoint": "GetRequestLocations"}).Info("Endpoint called in dev-nodb mode")
		c.IndentedJSON(http.StatusOK, black_kiwi_data_structs.MockRequestInfo)
		return
	}

	if from != "" && to != "" {
		fromTime, err := time.Parse("2006-01-02", from)
		if err != nil {
			log.WithFields(log.Fields{"from": from}).Warn("Unable to parse from time in GetRequestLocations")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse from time in GetRequestLocations"})
			return
		}

		toTime, err := time.Parse("2006-01-02", to)
		if err != nil {
			log.WithFields(log.Fields{"to": to}).Warn("Unable to parse to time in GetRequestLocations")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse to time in GetRequestLocations"})
			return
		}

		log.WithFields(log.Fields{"endpoint": "GetRequestLocations", "from": fromTime, "to": toTime}).Info("GetRequestLocations endpoint called")
		success, requestList = black_kiwi_admin_queries.GetRequestLocation(&fromTime, &toTime)
	} else {
		log.WithFields(log.Fields{"endpoint": "GetRequestLocations"}).Info("GetRequestLocations endpoint called")
		success, requestList = black_kiwi_admin_queries.GetRequestLocation(nil, nil)
	}

	if !success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting request locations"})
		return
	}

	if requestList == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No requests found"})
		return
	}

	c.IndentedJSON(http.StatusOK, *requestList)
}
