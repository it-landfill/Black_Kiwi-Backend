package black_kiwi_default

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ITLandfill/Black-Kiwi/structs/data_structs"
)

func GetPOIS(c *gin.Context) {

	// Allow CORS
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	c.IndentedJSON(http.StatusOK, black_kiwi_data_structs.MockPOIS)
}
