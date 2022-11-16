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
func GetPOIQuartieri(c *gin.Context) {	
	log.Info("GetPOIQuartieri endpoint called")

	var quartieriList *[]black_kiwi_data_structs.QuartiereInfo
	var result bool

	if os.Getenv("Black_Kiwi_ENV") == "dev-nodb" {
		log.WithFields(log.Fields{"endpoint": "GetPOIQuartieri"}).Info("Endpoint called in dev-nodb mode")
		quartieriList = &black_kiwi_data_structs.MockQuartiereInfo
		result = true
	} else {
		result, quartieriList = black_kiwi_admin_queries.GetPOIQuartieri()
	}

	if !result {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting quartieri info"})
		return
	}

	if quartieriList == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No quartieri found"})
		return
	}

	c.IndentedJSON(http.StatusOK, *quartieriList)
}
