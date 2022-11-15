package black_kiwi_admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/structs/data_structs"
)

// getAlbums responds with the list of all albums as JSON.
func GetRequestLocations(c *gin.Context) {
	from := c.DefaultQuery("from", "")
	to := c.DefaultQuery("to", "")

	log.WithFields(log.Fields{"endpoint": "GetRequestLocations", "from": from, "to": to}).Info("GetRequestLocations endpoint called")
	
	c.IndentedJSON(http.StatusOK, black_kiwi_data_structs.MockRequestInfo)
}
