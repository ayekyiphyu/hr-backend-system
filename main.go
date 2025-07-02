package main

import (
	"hr-backend-system/routes"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "hr-backend-system/docs"

	swaggerFiles "github.com/swaggo/files"
)

func main() {
	router := gin.Default()

	// Add Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Set up your actual routes
	routes.SetupRoutes(router)

	router.Run(":8080")
}
