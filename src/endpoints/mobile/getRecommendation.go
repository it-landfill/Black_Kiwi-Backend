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

	// No need to allow CORS on mobile endpoints

	category := c.DefaultQuery("category", "")
	minRank, err := strconv.ParseFloat(c.Query("minRank"), 32)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error while parsing limit in GetRecommendation")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error while parsing limit in GetRecommendation")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	latitude, err := strconv.ParseFloat(c.Query("latitude"), 64)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error while parsing latitude in GetRecommendation")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	longitude, err := strconv.ParseFloat(c.Query("longitude"), 64)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error while parsing longitude in GetRecommendation")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
