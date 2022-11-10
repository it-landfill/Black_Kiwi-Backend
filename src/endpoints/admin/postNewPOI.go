package black_kiwi_admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/structs/data_structs"
)

// getAlbums responds with the list of all albums as JSON.
func PostNewPOI(c *gin.Context) {

	// Allow CORS
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Get the POI from the request body
	body := black_kiwi_data_structs.PoiItem{}
	if err := c.BindJSON(&body); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error while parsing body in PostNewPOI")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.WithFields(log.Fields{"endpoint": "PostNewPOI", "body": body}).Info("New POI endpoint called")

	c.IndentedJSON(http.StatusOK, body)
}
