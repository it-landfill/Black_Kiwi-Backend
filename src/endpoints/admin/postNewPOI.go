package black_kiwi_admin

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/dbHandler/admin"
	"ITLandfill/Black-Kiwi/structs/data_structs"
)

// getAlbums responds with the list of all albums as JSON.
func PostNewPOI(c *gin.Context) {
	// Get the POI from the request body
	body := black_kiwi_data_structs.NewPoiItem{}
	if err := c.BindJSON(&body); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error while parsing body in PostNewPOI")
		
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing or malformed JSON body", "errorMessage": err.Error()})
		return
	}

	log.WithFields(log.Fields{"endpoint": "PostNewPOI", "body": body}).Info("New POI endpoint called")

	// Insert the POI into the database
	var poiID int

	if os.Getenv("Black_Kiwi_ENV") == "dev-nodb" {
		log.WithFields(log.Fields{"endpoint": "New POI"}).Info("Endpoint called in dev-nodb mode")
		poiID = 0
	} else {
		poiID = black_kiwi_admin_queries.AddPOI(body)
	}

	if poiID == -1 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while adding POI"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"poiID": poiID})
}
