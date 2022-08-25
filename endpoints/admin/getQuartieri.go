package black_kiwi_admin

import (
	"net/http"

	"ITLandfill/Black-Kiwi/structs"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func GetQuartieri(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, black_kiwi_structs.MockRequestInfo) //FIXME: Creare la query e struct e fare mock dell'output corretto
}