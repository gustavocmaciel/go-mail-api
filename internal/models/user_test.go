package models

import (
	"reflect"
	"testing"
)

// MockUserRepository is a mock implementation of UserRepository for testing purposes.
type MockUserRepository struct {
	Users []*User
	Err   error
}

func (m *MockUserRepository) Create(user *User) error {
	m.Users = append(m.Users, user)
	return m.Err
}

func (m *MockUserRepository) GetAllUsers() ([]*User, error) {
	return m.Users, m.Err
}

func TestCreateUser(t *testing.T) {
	// Create a new user with a mock repository
	repo := &MockUserRepository{}
	user := NewUser("test@example.com", "John", "Doe")
	// Call Create method
	err := repo.Create(user)
	// Assert that there is no error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Assert that the user was added to the repository
	if len(repo.Users) != 1 || !reflect.DeepEqual(repo.Users[0], user) {
		t.Errorf("User not added to the repository as expected.")
	}
}

func TestGetAllUsers(t *testing.T) {
	// Create a new user with a mock repository
	repo := &MockUserRepository{}
	user := NewUser("test@example.com", "John", "Doe")
	repo.Users = append(repo.Users, user)
	// Call GetAllUsers method
	users, err := repo.GetAllUsers()
	// Assert that there is no error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Assert that the returned users match the ones in the repository
	if len(users) != 1 || !reflect.DeepEqual(users[0], user) {
		t.Errorf("Returned users do not match the expected users.")
	}
}

func TestNewUser(t *testing.T) {
	// Create a new user
	user := NewUser("test@example.com", "John", "Doe")
	// Assert that the ID is not the zero value (indicating it was generated)
	if user.ID == "" {
		t.Errorf("NewUser did not generate a valid UUID for the user's ID.")
	}
	// Assert that the other fields are set correctly
	expectedUser := &User{
		ID:        user.ID,
		Email:     "test@example.com",
		FirstName: "John",
		LastName:  "Doe",
	}
	if !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("NewUser did not create the expected user.")
	}
}
