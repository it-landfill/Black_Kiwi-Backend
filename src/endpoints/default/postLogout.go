package black_kiwi_default

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"ITLandfill/Black-Kiwi/structs/auth_structs"
)

func PostLogout(c *gin.Context) {
	
	tokenStr := c.GetHeader("token")
	if tokenStr == "" {
		c.IndentedJSON(http.StatusOK, gin.H{"success": "No token found"})
		return
	}

	token := black_kiwi_auth_structs.GetToken(tokenStr)

	if token == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"success": "No token found"})
		return
	}

	if token != nil {
		(*token).DeleteToken()
	}

	c.IndentedJSON(http.StatusOK, gin.H{"success": "Logout completed"})
}
