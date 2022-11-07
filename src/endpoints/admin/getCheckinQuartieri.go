package black_kiwi_admin

import (
	"net/http"

	"github.com/gin-gonic/gin"


	"ITLandfill/Black-Kiwi/structs/data_structs"
)

// getAlbums responds with the list of all albums as JSON.
func GetCheckinQuartieri(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, black_kiwi_data_structs.MockQuartiereInfo)
}
