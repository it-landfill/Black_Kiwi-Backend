package main

import (
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"ITLandfill/Black-Kiwi/endpoints/admin"
	"ITLandfill/Black-Kiwi/endpoints/default"
	"ITLandfill/Black-Kiwi/endpoints/mobile"

	"ITLandfill/Black-Kiwi/dbHandler/utils"
	"ITLandfill/Black-Kiwi/dbHandler/handler"
)

func main() {

	if (os.Getenv("Black_Kiwi_ENV") != "") {
		log.WithFields(log.Fields{"Black_Kiwi_ENV": os.Getenv("Black_Kiwi_ENV")}).Info("Black_Kiwi_ENV is set")
	}

	if (os.Getenv("Black_Kiwi_ENV") != "dev" && os.Getenv("Black_Kiwi_ENV") != "dev-nodb") {
		gin.SetMode(gin.ReleaseMode)
	}

	if (os.Getenv("Black_Kiwi_ENV") != "dev-nodb") {
		// Init DB connection
		black_kiwi_db_utils.ConnPool = black_kiwi_db_handler.InitConnectionPool()
		defer black_kiwi_db_utils.ConnPool.Close()
	}
	

	// Generate a new router
    router := gin.Default()

	// Default
	router.GET("/", black_kiwi_default.GetRoot)
    router.GET("/getPOIS", black_kiwi_default.GetPOIS)
    router.GET("/getPOIS/:id", black_kiwi_default.GetPOI)

	// Admin
	router.GET("/getRequestLocations", black_kiwi_admin.GetRequestLocations)
    router.GET("/getPOIQuartieri", black_kiwi_admin.GetPOIQuartieri)
    router.GET("/getCheckinQuartieri", black_kiwi_admin.GetCheckinQuartieri)
	
	// Mobile
	router.GET("/getRecommendation", black_kiwi_mobile.GetRecommendation)


    router.Run("0.0.0.0:8080")
}