package handlers

import (
	"hr-backend-system/models"
	"hr-backend-system/storage"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary Get all users with pagination
// @Description Retrieve a paginated list of users
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Router /users [get]
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

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with name and email
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User creation request"
// @Success 201 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 409 {object} models.APIResponse
// @Router /users [post]
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

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Retrieve a user by their unique ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Router /users/{id} [get]
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

// UpdateUser godoc
// @Summary Update a user by ID
// @Description Update a user's name or email by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.UpdateUserRequest true "User update request"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Failure 409 {object} models.APIResponse
// @Router /users/{id} [put]
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

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
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

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Delete a user from the system by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Router /users/{id} [delete]
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
