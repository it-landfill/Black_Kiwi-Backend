package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/endpoints/admin"
	"ITLandfill/Black-Kiwi/endpoints/default"
	"ITLandfill/Black-Kiwi/endpoints/mobile"

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

	store := cookie.NewStore([]byte("secret")) // TODO: Change secret
	engine.Use(sessions.Sessions("sessiontoken", store))

	// Add CORS middleware to allow all origins
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	// Default
	engine.GET("/", black_kiwi_default.GetRoot)
	engine.GET("/getPOIS", black_kiwi_default.GetPOIS)
	engine.GET("/getPOIS/:id", black_kiwi_default.GetPOI)
	engine.POST("/login", black_kiwi_default.PostLogin)

	// Admin
	admin := engine.Group("/admin")
	admin.Use(AdminRequired)
	admin.GET("/getRequestLocations", black_kiwi_admin.GetRequestLocations)
	admin.GET("/getPOIQuartieri", black_kiwi_admin.GetPOIQuartieri)
	admin.GET("/getCheckinQuartieri", black_kiwi_admin.GetCheckinQuartieri)
	admin.POST("/newPOI", black_kiwi_admin.PostNewPOI)

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
	session := sessions.Default(c)
	role := session.Get("role")

	if role == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if role.(int8) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Wrong role"})
		return
	}

	c.Next()
}

// UserRequired is a simple middleware to check the session
func UserRequired(c *gin.Context) {
	session := sessions.Default(c)
	role := session.Get("role")

	if role == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if role.(int8) != 1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Wrong role"})
		return
	}

	c.Next()
}
