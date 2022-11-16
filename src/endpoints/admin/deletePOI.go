package black_kiwi_admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/structs/data_structs"
)

func DeletePOI(c *gin.Context) {
	poiID, err := strconv.Atoi(c.DefaultQuery("limit", "-1"))
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error while parsing poiID in DeletePOI")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	log.WithFields(log.Fields{"endpoint": "DeletePOI", "id": poiID}).Info("Delete POI endpoint called")

	c.IndentedJSON(http.StatusOK, black_kiwi_data_structs.MockPOIS[0])
}
