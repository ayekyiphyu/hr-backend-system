package handlers

import (
	"hr-backend-system/models"
	"hr-backend-system/storage"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetUsers gets all users with pagination
func GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	users := storage.GetAllUsers()

	start := (page - 1) * limit
	end := start + limit

	if start > len(users) {
		start = len(users)
	}
	if end > len(users) {
		end = len(users)
	}

	paginatedUsers := users[start:end]

	c.JSON(http.StatusOK, models.APIResponse{
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

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	var req models.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid request data",
			Error:   err.Error(),
		})
		return
	}

	// Check if email already exists
	if _, exists := storage.GetUserByEmail(req.Email); exists {
		c.JSON(http.StatusConflict, models.APIResponse{
			Success: false,
			Message: "User with this email already exists",
			Error:   "duplicate_email",
		})
		return
	}

	newUser := models.User{
		ID:        storage.GetNextUserID(),
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	storage.AddUser(newUser)

	c.JSON(http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "User created successfully",
		Data:    newUser,
	})
}

// GetUserByID gets a user by ID
func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid user ID",
			Error:   "invalid_id",
		})
		return
	}

	user, exists := storage.GetUserByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Message: "User not found",
			Error:   "user_not_found",
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "User retrieved successfully",
		Data:    user,
	})
}

// UpdateUser updates a user
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid user ID",
			Error:   "invalid_id",
		})
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid request data",
			Error:   err.Error(),
		})
		return
	}

	user, exists := storage.GetUserByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Message: "User not found",
			Error:   "user_not_found",
		})
		return
	}

	// Update fields if provided
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		// Check if new email already exists
		existingUser, exists := storage.GetUserByEmail(req.Email)
		if exists && existingUser.ID != id {
			c.JSON(http.StatusConflict, models.APIResponse{
				Success: false,
				Message: "Email already exists",
				Error:   "duplicate_email",
			})
			return
		}
		user.Email = req.Email
	}

	user.UpdatedAt = time.Now()

	storage.UpdateUser(id, user)

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "User updated successfully",
		Data:    user,
	})
}

// DeleteUser deletes a user
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid user ID",
			Error:   "invalid_id",
		})
		return
	}

	deletedUser, exists := storage.DeleteUser(id)
	if !exists {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Message: "User not found",
			Error:   "user_not_found",
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "User deleted successfully",
		Data: gin.H{
			"deleted_user": deletedUser,
		},
	})
}
