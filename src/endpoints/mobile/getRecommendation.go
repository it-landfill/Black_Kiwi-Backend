package black_kiwi_mobile

import (
	"net/http"
	"os"
	"strconv"

	black_kiwi_mobile_queries "ITLandfill/Black-Kiwi/dbHandler/mobile"
	"ITLandfill/Black-Kiwi/structs/data_structs"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetRecommendation(c *gin.Context) {

	category := c.DefaultQuery("category", "")

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error while parsing limit in GetRecommendation")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var minRank float64
	if rank := c.Query("minRank"); rank == "" {
		log.Warn("Missing minRank parameter in GetRecommendation")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing minRank parameter"})
		return
	} else {
		minRank, err = strconv.ParseFloat(rank, 32)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("Error while parsing limit in GetRecommendation")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}

	var latitude float64
	if lat := c.Query("latitude"); lat == "" {
		log.Warn("Missing latitude parameter in GetRecommendation")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing latitude parameter"})
		return
	} else {
		latitude, err = strconv.ParseFloat(lat, 32)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("Error while parsing latitude in GetRecommendation")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}

	var longitude float64
	if long := c.Query("longitude"); long == "" {
		log.Warn("Missing longitude parameter in GetRecommendation")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing longitude parameter"})
		return
	} else {
		longitude, err = strconv.ParseFloat(long, 32)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("Error while parsing longitude in GetRecommendation")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
	
	log.WithFields(log.Fields{
		"category":  category,
		"minRank":   minRank,
		"limit":     limit,
		"latitude":  latitude,
		"longitude": longitude,
	}).Info("New POI reccomendation requested")

	var res []black_kiwi_data_structs.PoiItem
	var success bool
	if os.Getenv("Black_Kiwi_ENV") != "dev-nodb" {
		success, res = black_kiwi_mobile_queries.GetRecommendation(minRank, latitude, longitude, category, limit)
	} else {
		res = black_kiwi_data_structs.MockPOIS
		success = true
	}

	if !success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting recommendation"})
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}
