package black_kiwi_admin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/dbHandler/admin"
)

// getAlbums responds with the list of all albums as JSON.
func GetRequestLocations(c *gin.Context) {
	from := c.DefaultQuery("from", "")
	to := c.DefaultQuery("to", "")

	log.WithFields(log.Fields{"endpoint": "GetRequestLocations", "from": from, "to": to}).Info("GetRequestLocations endpoint called")
	st, _ := time.Parse(time.RFC3339, "2022-10-01")
	et, _ := time.Parse(time.RFC3339, "2022-11-01")
	_, res := black_kiwi_admin_queries.GetRequestLocation(&st, &et)

	c.IndentedJSON(http.StatusOK, res)
}
