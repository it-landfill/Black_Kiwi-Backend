package black_kiwi_default

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/dbHandler/default"
	"ITLandfill/Black-Kiwi/structs/auth_structs"
)


func PostLogin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	log.WithFields(log.Fields{
		"endpoint": "PostLogin",
		"username": username,
		"password": password,
		}).Info("Endpoint called")

	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if os.Getenv("Black_Kiwi_ENV") == "dev-nodb" {
		log.WithFields(log.Fields{"endpoint": "PostLogin"}).Info("Endpoint called in dev-nodb mode")
		c.IndentedJSON(http.StatusOK, black_kiwi_auth_structs.MockUsers[0])
		return
	}
	success, user := black_kiwi_login_queries.GetUser(username, password)
	
	if !success {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
		return
	}

	//TODO: Generate and handle token
	user.Token = "I am a test token"

	c.IndentedJSON(http.StatusOK, user)
}