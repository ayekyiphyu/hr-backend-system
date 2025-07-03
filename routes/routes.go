package routes

import (
	"hr-backend-system/handlers"
	"hr-backend-system/middleware"
	"hr-backend-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all routes
func SetupRoutes(router *gin.Engine) {
	// Middleware
	router.Use(middleware.SetupCORS())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Root endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.APIResponse{
			Success: true,
			Message: "Welcome to Go REST API",
			Data: gin.H{
				"version": "1.0.0",
				"endpoints": gin.H{
					"health": "/api/v1/health",
					"users":  "/api/v1/users",
					"auth":   "/api/v1/auth/login",
				},
			},
		})
	})

	// API routes
	api := router.Group("/api/v1")
	{
		// Health check
		api.GET("/health", handlers.HealthCheck)

		// User routes
		users := api.Group("/users")
		{
			users.GET("", handlers.GetUsers)
			users.POST("", handlers.CreateUser)
			users.GET("/:id", handlers.GetUserByID)
			users.PUT("/:id", handlers.UpdateUser)
			users.DELETE("/:id", handlers.DeleteUser)
		}
	}
}
