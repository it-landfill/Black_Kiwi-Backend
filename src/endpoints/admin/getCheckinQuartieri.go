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
func GetCheckinQuartieri(c *gin.Context) {
	log.Info("GetCheckinQuartieri endpoint called")
	
	var featureCollection string
	var result bool

	if os.Getenv("Black_Kiwi_ENV") == "dev-nodb" {
		log.WithFields(log.Fields{"endpoint": "GetCheckinQuartieri"}).Info("Endpoint called in dev-nodb mode")
		featureCollection = black_kiwi_data_structs.MockQuartiereInfo
		result = true
	} else {
		result, featureCollection = black_kiwi_admin_queries.GetCheckinQuartieri()
	}

	if !result {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting quartieri info"})
		return
	}

	if featureCollection == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No quartieri found"})
		return
	}

	c.IndentedJSON(http.StatusOK, featureCollection)
}
	