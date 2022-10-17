package black_kiwi_mobile

import (
	"net/http"
	"strconv"

	"ITLandfill/Black-Kiwi/dbHandler/mobile"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)


func GetRecommendation(c *gin.Context) {
    
	category := c.DefaultQuery("category", "")
	minRank, err := strconv.Atoi(c.Query("minRank"))
	if err != nil {
		log.WithFields(log.Fields{"error":err,}).Error("Error while parsing limit in GetRecommendation")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		log.WithFields(log.Fields{"error":err,}).Error("Error while parsing limit in GetRecommendation")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	latitude, err := strconv.ParseFloat(c.Query("latitude"), 64)
	if err != nil {
		log.WithFields(log.Fields{"error":err,}).Error("Error while parsing latitude in GetRecommendation")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	longitude, err := strconv.ParseFloat(c.Query("longitude"), 64)
	if err != nil {
		log.WithFields(log.Fields{"error":err,}).Error("Error while parsing longitude in GetRecommendation")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	log.WithFields(log.Fields{
		"category": category,
		"minRank":   minRank,
		"limit": limit,
		"latitude": latitude,
		"longitude": longitude,
	  }).Info("New POI reccomendation requested")

	res := black_kiwi_mobile_queries.GetRecommendation(minRank, latitude, longitude, category, limit)

    c.IndentedJSON(http.StatusOK, res)
}