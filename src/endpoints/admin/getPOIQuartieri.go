package black_kiwi_admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/structs/data_structs"
)

// getAlbums responds with the list of all albums as JSON.
func GetPOIQuartieri(c *gin.Context) {	
	log.Info("GetPOIQuartieri endpoint called")

	c.IndentedJSON(http.StatusOK, black_kiwi_data_structs.MockQuartiereInfo)
}
