package handlers

import (
	"hr-backend-system/models"
	"hr-backend-system/storage"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	// Set maximum limit to prevent abuse
	if limit > 100 {
		limit = 100
	}

	users := storage.GetAllUsers()
	total := len(users)

	// Calculate pagination bounds
	start := (page - 1) * limit
	if start >= total {
		// Return empty result if page is beyond available data
		c.JSON(http.StatusOK, models.APIResponse{
			Success: true,
			Message: "Users retrieved successfully",
			Data: gin.H{
				"users": []models.User{},
				"pagination": gin.H{
					"page":        page,
					"limit":       limit,
					"total":       total,
					"total_pages": (total + limit - 1) / limit,
					"has_more":    false,
				},
			},
		})
		return
	}

	end := start + limit
	if end > total {
		end = total
	}

	paginatedUsers := users[start:end]

	// Remove sensitive data from response
	sanitizedUsers := make([]models.User, len(paginatedUsers))
	for i, user := range paginatedUsers {
		sanitizedUsers[i] = user
		sanitizedUsers[i].Password = "" // Don't expose passwords
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Users retrieved successfully",
		Data: gin.H{
			"users": sanitizedUsers,
			"pagination": gin.H{
				"page":        page,
				"limit":       limit,
				"total":       total,
				"total_pages": (total + limit - 1) / limit,
				"has_more":    end < total,
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

	// Validate required fields
	if strings.TrimSpace(req.Name) == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Name is required",
			Error:   "missing_name",
		})
		return
	}

	if strings.TrimSpace(req.Email) == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Email is required",
			Error:   "missing_email",
		})
		return
	}

	if strings.TrimSpace(req.Password) == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Password is required",
			Error:   "missing_password",
		})
		return
	}

	// Basic email validation
	if !strings.Contains(req.Email, "@") {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid email format",
			Error:   "invalid_email",
		})
		return
	}

	// Check if user already exists
	if _, exists := storage.GetUserByEmail(req.Email); exists {
		c.JSON(http.StatusConflict, models.APIResponse{
			Success: false,
			Message: "User with this email already exists",
			Error:   "duplicate_email",
		})
		return
	}

	// Hash password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to process password",
			Error:   "password_hash_error",
		})
		return
	}

	newUser := models.User{
		ID:        storage.GetNextUserID(),
		Name:      strings.TrimSpace(req.Name),
		Email:     strings.ToLower(strings.TrimSpace(req.Email)),
		Type:      req.Type,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	storage.AddUser(newUser)

	// Don't return password in response
	newUser.Password = ""

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
	if err != nil || id < 1 {
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

	// Don't expose password
	user.Password = ""

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
	if err != nil || id < 1 {
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

	// Update name if provided
	if req.Name != "" {
		user.Name = strings.TrimSpace(req.Name)
	}

	// Update email if provided
	if req.Email != "" {
		email := strings.ToLower(strings.TrimSpace(req.Email))

		// Basic email validation
		if !strings.Contains(email, "@") {
			c.JSON(http.StatusBadRequest, models.APIResponse{
				Success: false,
				Message: "Invalid email format",
				Error:   "invalid_email",
			})
			return
		}

		// Check for duplicate email
		existingUser, exists := storage.GetUserByEmail(email)
		if exists && existingUser.ID != id {
			c.JSON(http.StatusConflict, models.APIResponse{
				Success: false,
				Message: "Email already exists",
				Error:   "duplicate_email",
			})
			return
		}
		user.Email = email
	}

	// Update password if provided
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse{
				Success: false,
				Message: "Failed to process password",
				Error:   "password_hash_error",
			})
			return
		}
		user.Password = string(hashedPassword)
	}

	// Update type if provided
	if req.Type != "" {
		user.Type = req.Type
	}

	user.UpdatedAt = time.Now()

	storage.UpdateUser(id, user)

	// Don't return password in response
	user.Password = ""

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
	if err != nil || id < 1 {
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

	// Don't return password even for deleted user
	deletedUser.Password = ""

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "User deleted successfully",
		Data: gin.H{
			"deleted_user": deletedUser,
		},
	})
}
