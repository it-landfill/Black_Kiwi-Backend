package main

import (
    "github.com/gin-gonic/gin"
	
	"ITLandfill/Black-Kiwi/endpoints/default"
	"ITLandfill/Black-Kiwi/endpoints/admin"
	"ITLandfill/Black-Kiwi/endpoints/mobile"
)

func main() {
    router := gin.Default()

	// Default
    router.GET("/getPOIS", black_kiwi_default.GetPOIS)
    router.GET("/getPOI/:id", black_kiwi_default.GetPOI)

	// Admin
	router.GET("/getRequestLocations", black_kiwi_admin.GetRequestLocations)
    router.GET("/getQuartieri", black_kiwi_admin.GetQuartieri)
	
	// Mobile
	router.GET("/getRecommendation", black_kiwi_mobile.GetRecommendation)


    router.Run("localhost:8080")
}