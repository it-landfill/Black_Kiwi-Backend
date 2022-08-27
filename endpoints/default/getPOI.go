package black_kiwi_default

import (
	"net/http"
	"strconv"

	"ITLandfill/Black-Kiwi/structs"

	"github.com/gin-gonic/gin"
)


func GetPOI(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"),10,0)

	if err != nil {
		println("Oh shit")
	}


    for _, a := range black_kiwi_structs.MockPOIS {
        if int64(a.Id) == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "POI not found"})
}