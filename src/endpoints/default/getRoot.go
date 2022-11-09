package black_kiwi_default

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func GetRoot(c *gin.Context) {

	// Allow CORS
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
    c.String(http.StatusOK, "What are you doing here?")
}