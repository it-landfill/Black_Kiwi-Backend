package black_kiwi_admin

import (
	"net/http"
	"strconv"
	"os"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/dbHandler/admin"
)

func DeletePOI(c *gin.Context) {
	poiID, err := strconv.Atoi(c.DefaultQuery("poiID", "-1"))
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error while parsing poiID in DeletePOI")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	log.WithFields(log.Fields{"endpoint": "DeletePOI", "id": poiID}).Info("Delete POI endpoint called")

	var success bool
	if os.Getenv("Black_Kiwi_ENV") == "dev-nodb" {
		success = true
	} else {
		success = black_kiwi_admin_queries.DeletePOI(poiID)
	}

	if !success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting POI"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"success": true})
}
