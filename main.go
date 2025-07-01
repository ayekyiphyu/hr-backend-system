package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// User represents a user in our system
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Age       int       `json:"age" binding:"min=1,max=120"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserRequest for creating users
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"min=1,max=120"`
}

// UpdateUserRequest for updating users
type UpdateUserRequest struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Age   int    `json:"age,omitempty"`
}

// APIResponse standard response format
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// In-memory storage (use database in production)
var (
	users       []User
	userCounter int = 1
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Middleware for logging
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// API routes
	api := router.Group("/api/v1")
	{
		// Health check
		api.GET("/health", healthCheck)

		// User routes
		users := api.Group("/users")
		{
			users.GET("", getUsers)
			users.POST("", createUser)
			users.GET("/:id", getUserByID)
			users.PUT("/:id", updateUser)
			users.DELETE("/:id", deleteUser)
		}
	}

	// Root endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, APIResponse{
			Success: true,
			Message: "Welcome to Go REST API",
			Data: gin.H{
				"version": "1.0.0",
				"endpoints": gin.H{
					"health": "/api/v1/health",
					"users":  "/api/v1/users",
				},
			},
		})
	})

	// Start server
	router.Run(":8080")
}

// Health check endpoint
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "API is healthy",
		Data: gin.H{
			"timestamp": time.Now(),
			"status":    "up",
		},
	})
}

// Get all users
func getUsers(c *gin.Context) {
	// Query parameters for pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Simple pagination
	start := (page - 1) * limit
	end := start + limit

	if start > len(users) {
		start = len(users)
	}
	if end > len(users) {
		end = len(users)
	}

	paginatedUsers := users[start:end]

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Users retrieved successfully",
		Data: gin.H{
			"users": paginatedUsers,
			"pagination": gin.H{
				"page":     page,
				"limit":    limit,
				"total":    len(users),
				"has_more": end < len(users),
			},
		},
	})
}

// Create new user
func createUser(c *gin.Context) {
	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Invalid request data",
			Error:   err.Error(),
		})
		return
	}

	// Check if email already exists
	for _, user := range users {
		if user.Email == req.Email {
			c.JSON(http.StatusConflict, APIResponse{
				Success: false,
				Message: "User with this email already exists",
				Error:   "duplicate_email",
			})
			return
		}
	}

	newUser := User{
		ID:        userCounter,
		Name:      req.Name,
		Email:     req.Email,
		Age:       req.Age,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	users = append(users, newUser)
	userCounter++

	c.JSON(http.StatusCreated, APIResponse{
		Success: true,
		Message: "User created successfully",
		Data:    newUser,
	})
}

// Get user by ID
func getUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Invalid user ID",
			Error:   "invalid_id",
		})
		return
	}

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, APIResponse{
				Success: true,
				Message: "User retrieved successfully",
				Data:    user,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Success: false,
		Message: "User not found",
		Error:   "user_not_found",
	})
}

// Update user
func updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Invalid user ID",
			Error:   "invalid_id",
		})
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Invalid request data",
			Error:   err.Error(),
		})
		return
	}

	for i, user := range users {
		if user.ID == id {
			// Update fields if provided
			if req.Name != "" {
				users[i].Name = req.Name
			}
			if req.Email != "" {
				// Check if new email already exists
				for j, existingUser := range users {
					if j != i && existingUser.Email == req.Email {
						c.JSON(http.StatusConflict, APIResponse{
							Success: false,
							Message: "Email already exists",
							Error:   "duplicate_email",
						})
						return
					}
				}
				users[i].Email = req.Email
			}
			if req.Age > 0 {
				users[i].Age = req.Age
			}
			users[i].UpdatedAt = time.Now()

			c.JSON(http.StatusOK, APIResponse{
				Success: true,
				Message: "User updated successfully",
				Data:    users[i],
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Success: false,
		Message: "User not found",
		Error:   "user_not_found",
	})
}

// Delete user
func deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Invalid user ID",
			Error:   "invalid_id",
		})
		return
	}

	for i, user := range users {
		if user.ID == id {
			// Remove user from slice
			users = append(users[:i], users[i+1:]...)

			c.JSON(http.StatusOK, APIResponse{
				Success: true,
				Message: "User deleted successfully",
				Data: gin.H{
					"deleted_user": user,
				},
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Success: false,
		Message: "User not found",
		Error:   "user_not_found",
	})
}
