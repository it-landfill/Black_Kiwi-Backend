package black_kiwi_admin

import (
	"net/http"

	"ITLandfill/Black-Kiwi/structs"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func GetCheckinQuartieri(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, black_kiwi_structs.MockQuartiereInfo)
}