package main

import (
	"github.com/gin-gonic/gin"

	"ITLandfill/Black-Kiwi/endpoints/admin"
	"ITLandfill/Black-Kiwi/endpoints/default"
	"ITLandfill/Black-Kiwi/endpoints/mobile"
)

func main() {
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