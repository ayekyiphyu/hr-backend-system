package models

import "time"

// User represents a user in our system
type User struct {
	ID        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"John Doe"`
	Email     string    `json:"email" example:"john@example.com"`
	Type      string    `json:"type" example:"jobseeker"`
	Password  string    `json:"-"` // Do not expose in JSON responses
	CreatedAt time.Time `json:"created_at" example:"2025-07-02T15:04:05Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2025-07-02T15:04:05Z"`
}

// LoginRequest represents the request payload for user login/authentication
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"yourpassword"`
}

// CreateUserRequest represents the request payload for creating a user
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Type     string `json:"type" binding:"required,oneof=admin-staff organization jobseeker" example:"jobseeker"`
	Password string `json:"password" binding:"required,min=8,max=128" example:"securepassword123"`
}

// UpdateUserRequest represents the request payload for updating a user
type UpdateUserRequest struct {
	Name     string `json:"name,omitempty" binding:"omitempty,min=2,max=100" example:"Jane Doe"`
	Email    string `json:"email,omitempty" binding:"omitempty,email" example:"jane@example.com"`
	Type     string `json:"type,omitempty" binding:"omitempty,oneof=admin-staff organization jobseeker" example:"organization"`
	Password string `json:"password,omitempty" binding:"omitempty,min=8,max=128" example:"newsecurepassword456"`
}

// RegisterRequest represents the request payload for user registration (extended version)
type RegisterRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=100" example:"John Doe"`
	Email       string `json:"email" binding:"required,email" example:"john@example.com"`
	PhoneNumber string `json:"phone_number" binding:"required,min=10,max=15" example:"080-1234-5678"`
	Type        string `json:"type" binding:"required,oneof=admin-staff organization jobseeker" example:"jobseeker"`
	Password    string `json:"password" binding:"required,min=8,max=128" example:"securepassword123"`
}

// UserResponse represents the user data returned in API responses (without sensitive info)
type UserResponse struct {
	ID        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"John Doe"`
	Email     string    `json:"email" example:"john@example.com"`
	Type      string    `json:"type" example:"jobseeker"`
	CreatedAt time.Time `json:"created_at" example:"2025-07-02T15:04:05Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2025-07-02T15:04:05Z"`
}

// ChangePasswordRequest represents the request payload for changing password
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required,min=8" example:"oldpassword123"`
	NewPassword     string `json:"new_password" binding:"required,min=8,max=128" example:"newpassword456"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=NewPassword" example:"newpassword456"`
}

// UserListResponse represents paginated user list response
type UserListResponse struct {
	Users      []UserResponse `json:"users"`
	Pagination PaginationInfo `json:"pagination"`
}

// PaginationInfo represents pagination metadata
type PaginationInfo struct {
	Page       int  `json:"page" example:"1"`
	Limit      int  `json:"limit" example:"10"`
	Total      int  `json:"total" example:"100"`
	TotalPages int  `json:"total_pages" example:"10"`
	HasMore    bool `json:"has_more" example:"true"`
}

// Helper method to convert User to UserResponse
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Type:      u.Type,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
