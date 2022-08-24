package black_kiwi

import (
    "net/http"

    "github.com/gin-gonic/gin"
	"ITLandfill/Black-Kiwi/structs"
)

func GetPOIS(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, black_kiwi.MockPOIS)
}