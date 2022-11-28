package black_kiwi_admin

import (
	"net/http"
	"os"


	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/structs/data_structs"
	"ITLandfill/Black-Kiwi/dbHandler/admin"
)

// getAlbums responds with the list of all albums as JSON.
func EditPOI(c *gin.Context) {
	// Get the POI from the request body
	body := black_kiwi_data_structs.PoiItem{}
	if err := c.BindJSON(&body); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error while parsing body in EditPOI")
		
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing or malformed JSON body", "errorMessage": err.Error()})
		return
	}

	log.WithFields(log.Fields{"endpoint": "EditPOI", "new POI": body}).Info("Edit POI endpoint called")

	var success bool
	if os.Getenv("Black_Kiwi_ENV") == "dev-nodb" {
		success = true
	} else {
		success = black_kiwi_admin_queries.EditPOI(body)
	}

	if !success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while editing POI"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"success": true})
}
