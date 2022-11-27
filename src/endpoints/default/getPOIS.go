package black_kiwi_default

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/dbHandler/default"
	"ITLandfill/Black-Kiwi/structs/data_structs"
)

func GetPOIS(c *gin.Context) {

	log.Info("GetPOIS endpoint called")

	var poiList *[]black_kiwi_data_structs.PoiItem
	var result bool

	if os.Getenv("Black_Kiwi_ENV") == "dev-nodb" {
		log.WithFields(log.Fields{"endpoint": "GetPOIS"}).Info("Endpoint called in dev-nodb mode")
		poiList = &black_kiwi_data_structs.MockPOIS
		result = true
	} else {
		result, poiList = black_kiwi_default_queries.GetPOIS()
	}

	if !result {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting user"})
		return
	}

	if poiList == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No Pois found"})
		return
	}

	c.IndentedJSON(http.StatusOK, *poiList)
}
