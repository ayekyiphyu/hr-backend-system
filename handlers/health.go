package handlers

import (
	"hr-backend-system/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthCheck endpoint
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "API is healthy",
		Data: gin.H{
			"timestamp": time.Now(),
			"status":    "up",
		},
	})
}
