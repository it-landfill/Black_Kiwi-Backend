package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/endpoints/admin"
	"ITLandfill/Black-Kiwi/endpoints/default"
	"ITLandfill/Black-Kiwi/endpoints/mobile"

	"ITLandfill/Black-Kiwi/structs/auth_structs"

	"ITLandfill/Black-Kiwi/dbHandler/handler"
	"ITLandfill/Black-Kiwi/dbHandler/utils"
)

func main() {

	if os.Getenv("Black_Kiwi_ENV") != "" {
		log.WithFields(log.Fields{"Black_Kiwi_ENV": os.Getenv("Black_Kiwi_ENV")}).Info("Black_Kiwi_ENV is set")
	}

	if os.Getenv("Black_Kiwi_ENV") != "dev" && os.Getenv("Black_Kiwi_ENV") != "dev-nodb" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		log.WithFields(log.Fields{"Black_Kiwi_ENV": os.Getenv("Black_Kiwi_ENV")}).Warn("Black_Kiwi_ENV is set to debug mode")
	}

	if os.Getenv("Black_Kiwi_ENV") != "dev-nodb" {
		// Init DB connection
		black_kiwi_db_utils.ConnPool = black_kiwi_db_handler.InitConnectionPool()
		defer black_kiwi_db_utils.ConnPool.Close()
	}

	// Initialize the token store
	black_kiwi_auth_structs.InitTokenArr()

	// Generate a new router
	router := createEngine()

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}

func createEngine() *gin.Engine {
	// gin.Default() has already added Logger and Recovery middleware.
	// gin.New() is used to create a new engine without any middleware attached.
	engine := gin.Default()

	// Add CORS middleware to allow all origins
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "DELETE"},
		AllowHeaders: []string{"Content-Type", "X-API-KEY"},
	}))

	// Default
	engine.GET("/", black_kiwi_default.GetRoot)
	engine.GET("/pois", black_kiwi_default.GetPOIS)
	engine.GET("/pois/:id", black_kiwi_default.GetPOI)
	engine.POST("/login", black_kiwi_default.PostLogin)
	engine.POST("/logout", black_kiwi_default.PostLogout)

	// Admin
	admin := engine.Group("/admin")
	admin.Use(AdminRequired)
	admin.GET("/getRequestLocations", black_kiwi_admin.GetRequestLocations)
	admin.GET("/getPOIQuartieri", black_kiwi_admin.GetPOIQuartieri)
	admin.GET("/getCheckinQuartieri", black_kiwi_admin.GetCheckinQuartieri)
	admin.POST("/newPOI", black_kiwi_admin.PostNewPOI)
	admin.POST("/editPOI", black_kiwi_admin.EditPOI)
	admin.DELETE("/deletePOI", black_kiwi_admin.DeletePOI)

	// Mobile
	mobile := engine.Group("/mobile")
	mobile.Use(UserRequired)
	mobile.GET("/getRecommendation", black_kiwi_mobile.GetRecommendation)
	mobile.GET("/testAuth", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Ok")
	})

	return engine
}

// AdminRequired is a simple middleware to check the session
func AdminRequired(c *gin.Context) {
	
	tokenStr := c.GetHeader("X-API-KEY")

	if tokenStr == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}

	token := black_kiwi_auth_structs.GetToken(tokenStr)
	
	if token == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if token != nil && (*token).IsExpired() {
		(*token).DeleteToken()
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Expired token"})
		return
	}

	if token != nil && !(*token).IsAdmin() {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not admin"})
		return
	}

	c.Next()
}

// UserRequired is a simple middleware to check the session
func UserRequired(c *gin.Context) {
	tokenStr := c.GetHeader("token")

	if tokenStr == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}

	token := black_kiwi_auth_structs.GetToken(tokenStr)
	
	if token == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if token != nil && (*token).IsExpired() {
		(*token).DeleteToken()
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Expired token"})
		return
	}

	if token != nil && !(*token).IsUser() {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not user"})
		return
	}

	c.Next()
}
