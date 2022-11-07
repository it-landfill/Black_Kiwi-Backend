package black_kiwi_default

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/dbHandler/default"
)


func PostLogin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	log.WithFields(log.Fields{
		"endpoint": "PostLogin",
		"username": username,
		"password": password,
		}).Info("Endpoint called")

	success, user := black_kiwi_login_queries.GetUser(username, password)
	
	if !success {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
		return
	}

	//TODO: Generate and handle token
	user.Token = "I am a test token"

	c.IndentedJSON(http.StatusOK, user)
}