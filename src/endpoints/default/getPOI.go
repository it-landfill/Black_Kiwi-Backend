package black_kiwi_default

import (
	"net/http"
	"strconv"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/dbHandler/default"
	"ITLandfill/Black-Kiwi/structs/data_structs"
)

func GetPOI(c *gin.Context) {
	poiIDStr := c.Param("id")
	poiID, err := strconv.ParseInt(poiIDStr, 10, 0)
	if err != nil {
		log.WithField("poiID", poiIDStr).Warn("Unable to parse POI ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse POI ID"})
		return
	}

	log.WithField("poiID", poiID).Info("GetPOI endpoint called")

	var poi *black_kiwi_data_structs.PoiItem
	var result bool

	if os.Getenv("Black_Kiwi_ENV") == "dev-nodb" {
		log.WithFields(log.Fields{"endpoint": "GetPOIS"}).Info("Endpoint called in dev-nodb mode")
		poi = &black_kiwi_data_structs.MockPOIS[0]
		result = true
	} else {
		result, poi = black_kiwi_default_queries.GetPOI(int(poiID))
	}

	if !result {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting user"})
		return
	}

	if poi == nil {
		c.IndentedJSON(http.StatusOK, "{}")
		return
	}

	c.IndentedJSON(http.StatusOK, *poi)
}
