package black_kiwi

import (
	"net/http"
	"strconv"

	"ITLandfill/Black-Kiwi/structs"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func GetPOI(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"),10,0)

	if err != nil {
		println("Oh shit")
	}

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range black_kiwi.MockPOIS {
        if int64(a.Id) == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}