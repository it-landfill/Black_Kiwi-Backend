package black_kiwi_default

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func GetRoot(c *gin.Context) {
    c.String(http.StatusOK, "What are you doing here?")
}