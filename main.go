package main

import (
	"url-shortener/src/config"
	"url-shortener/src/db"
	"url-shortener/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.SetConfig()

	db.ConnectToRedis()

	// Uncomment gin.SetMode when in production
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	routes.AddAppRoutes(router)

	port := config.ConfigVariables.Port
	router.Run(":" + port)
}
