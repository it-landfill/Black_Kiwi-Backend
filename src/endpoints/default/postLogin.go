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

	// Allow CORS
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	username := c.PostForm("username")
	password := c.PostForm("password")

	log.WithFields(log.Fields{
		"endpoint": "PostLogin",
		"username": username,
		"password": password,
		}).Info("Endpoint called")

	var success bool
	var user black_kiwi_auth_structs.User

	if os.Getenv("Black_Kiwi_ENV") == "dev-nodb" {
		log.WithFields(log.Fields{"endpoint": "PostLogin"}).Info("Endpoint called in dev-nodb mode")
		success = true
		user = black_kiwi_auth_structs.MockUsers[0]
	}
	success, user = black_kiwi_login_queries.GetUser(username, password)
	
	if !success {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
		return
	}

	//TODO: Generate and handle token
	user.Token = "I am a test token"

	c.IndentedJSON(http.StatusOK, user)
}