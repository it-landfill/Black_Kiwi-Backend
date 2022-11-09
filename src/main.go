package main

import (
	"os"

	"github.com/gin-contrib/cors"
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
	engine.GET("/getRequestLocations", black_kiwi_admin.GetRequestLocations)
	engine.GET("/getPOIQuartieri", black_kiwi_admin.GetPOIQuartieri)
	engine.GET("/getCheckinQuartieri", black_kiwi_admin.GetCheckinQuartieri)

	// Mobile
	engine.GET("/getRecommendation", black_kiwi_mobile.GetRecommendation)

	return engine
}