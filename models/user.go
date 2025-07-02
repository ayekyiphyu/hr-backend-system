package models

import "time"

// User represents a user in our system
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserRequest struct {
	Name        string `json:"name" example:"John" binding:"required" `
	Email       string `json:"emil" example:"john@example.com" binding:"required"`
	PhoneNumber string `json:"phone_number" example:"080-1234-5678" binding:"required"`
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
