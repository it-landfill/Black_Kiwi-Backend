package black_kiwi_mobile

import (
	"net/http"

	"ITLandfill/Black-Kiwi/structs"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// getAlbums responds with the list of all albums as JSON.
func GetRecommendation(c *gin.Context) {
    
	category := black_kiwi_structs.Categories(c.Query("category"))
	minRank := c.DefaultQuery("minRank", "0")
	limit := c.DefaultQuery("limit", "5")
	latitude := c.Query("latitude")
	longitude := c.Query("longitude")

	log.WithFields(log.Fields{
		"category": category,
		"minRank":   minRank,
		"limit": limit,
		"latitude": latitude,
		"longitude": longitude,
	  }).Info("New POI reccomendation requested")

    c.IndentedJSON(http.StatusOK, black_kiwi_structs.MockPOIS[0:2])
}