package models

import "github.com/google/uuid"

// UserRepository defines the methods that a repository for managing user data should implement.
type UserRepository interface {
	// Create inserts a new user into the repository.
	Create(user *User) error

	// GetAllUsers retrieves a list of all users from the repository.
	GetAllUsers() ([]*User, error)
}

// User represents a user entity with basic information.
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func NewUser(email string, firstName string, lastName string) *User {
	return &User{
		ID:        uuid.New().String(),
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
}
