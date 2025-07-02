package models

import "time"

// User represents a user in our system
type User struct {
	ID        int       `json:"id" example:"1"`
	Name      string    `json:"name" binding:"required" example:"John Doe"`
	Email     string    `json:"email" binding:"required,email" example:"john@example.com"`
	CreatedAt time.Time `json:"created_at" example:"2025-07-02T15:04:05Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2025-07-02T15:04:05Z"`
}

// GetUserRequest represents the request payload for getting a user
type GetUserRequest struct {
	Name        string `json:"name" example:"John" binding:"required"`
	Email       string `json:"email" example:"john@example.com" binding:"required,email"`
	PhoneNumber string `json:"phone_number" example:"080-1234-5678" binding:"required"`
}

// CreateUserRequest represents the request payload for creating a user
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required" example:"John Doe"`
	Email string `json:"email" binding:"required,email" example:"john@example.com"`
	Age   int    `json:"age" binding:"min=1,max=120" example:"30"`
}

// UpdateUserRequest represents the request payload for updating a user
type UpdateUserRequest struct {
	Name  string `json:"name,omitempty" example:"Jane Doe"`
	Email string `json:"email,omitempty" example:"jane@example.com"`
	Age   int    `json:"age,omitempty" example:"35"`
}
