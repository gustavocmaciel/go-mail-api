package models

import "github.com/google/uuid"

// UserRepository defines the methods that a repository for managing user data should implement.
type UserRepository interface {
	// Inserts a new user into the repository.
	// Returns an error if the operation fails.
	Create(user *User) error

	// Retrieves a list of all users from the repository.
	//
	// Returns:
	//   - A slice of pointers to User structs representing all users.
	//   - An error if the operation fails.
	GetAllUsers() ([]*User, error)
}

// User represents a user entity with basic information.
type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}

// NewUser creates a new User instance with the provided email, first name, and last name.
// It generates a new UUID for the user's ID.
//
// Parameters:
//   - email: The email address of the user.
//   - firstName: The first name of the user.
//   - lastName: The last name of the user.
//
// Returns:
//   - A pointer to the newly created User instance.
func NewUser(email string, firstName string, lastName string) *User {
	return &User{
		ID:        uuid.New(),
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
}
