package black_kiwi_default

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/dbHandler/default"
	"ITLandfill/Black-Kiwi/structs/auth_structs"
)

func PostLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	role, err := strconv.Atoi(c.DefaultPostForm("role", "-1"))
	if err != nil {
		log.WithFields(log.Fields{"error": err, "role": role}).Error("Error while parsing role in PostLogin")
		c.AbortWithError(http.StatusBadRequest, err)
	}

	log.WithFields(log.Fields{
		"endpoint": "PostLogin",
		"username": username,
		"password": password,
	}).Info("Endpoint called")

	var success bool
	var user *black_kiwi_auth_structs.User

	if os.Getenv("Black_Kiwi_ENV") == "dev-nodb" {
		log.WithFields(log.Fields{"endpoint": "PostLogin"}).Info("Endpoint called in dev-nodb mode")
		success = true
		user = &black_kiwi_auth_structs.MockUsers[0]
	} else {
		success, user = black_kiwi_login_queries.GetUser(username, password)
	}

	if !success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting user"})
		return
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong username or password"})
		return
	}

	if role != -1 && (*user).Role != int8(role) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong role"})
		return
	}

	black_kiwi_auth_structs.AddToken(user)

	c.IndentedJSON(http.StatusOK, *user)
}
