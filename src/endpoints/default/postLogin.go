package black_kiwi_default

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/structs/auth_structs"
)


func PostLogin(c *gin.Context) {
	log.WithFields(log.Fields{
		"endpoint": "PostLogin",
		"username": c.Param("username"),
		"password": c.Param("password"),
		}).Info("Endpoint called")

	user := black_kiwi_auth_structs.MockUsers[0]
	c.IndentedJSON(http.StatusOK, user)
}