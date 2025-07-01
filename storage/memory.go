package storage

import (
	"hr-backend-system/models"
	"sync"
)

var (
	users       []models.User
	userCounter int = 1
	mu          sync.RWMutex
)

// GetAllUsers returns all users
func GetAllUsers() []models.User {
	mu.RLock()
	defer mu.RUnlock()
	return users
}

// AddUser adds a new user
func AddUser(user models.User) {
	mu.Lock()
	defer mu.Unlock()
	users = append(users, user)
}

// GetUserByID returns a user by ID
func GetUserByID(id int) (models.User, bool) {
	mu.RLock()
	defer mu.RUnlock()
	for _, user := range users {
		if user.ID == id {
			return user, true
		}
	}
	return models.User{}, false
}

// GetUserByEmail returns a user by email
func GetUserByEmail(email string) (models.User, bool) {
	mu.RLock()
	defer mu.RUnlock()
	for _, user := range users {
		if user.Email == email {
			return user, true
		}
	}
	return models.User{}, false
}

// UpdateUser updates a user
func UpdateUser(id int, updatedUser models.User) bool {
	mu.Lock()
	defer mu.Unlock()
	for i, user := range users {
		if user.ID == id {
			users[i] = updatedUser
			return true
		}
	}
	return false
}

// DeleteUser deletes a user
func DeleteUser(id int) (models.User, bool) {
	mu.Lock()
	defer mu.Unlock()
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return user, true
		}
	}
	return models.User{}, false
}

// GetNextUserID returns the next available user ID
func GetNextUserID() int {
	mu.Lock()
	defer mu.Unlock()
	id := userCounter
	userCounter++
	return id
}
