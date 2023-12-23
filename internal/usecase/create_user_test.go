package usecase

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/gustavocmaciel/go-mail-api/internal/models"
	"github.com/stretchr/testify/assert"
)

// MockUserRepository is a mock implementation of UserRepository for testing purposes.
type MockUserRepository struct {
	CreateFunc      func(user *models.User) error
	GetAllUsersFunc func() ([]*models.User, error)
}

func (m *MockUserRepository) Create(user *models.User) error {
	return m.CreateFunc(user)
}

func (m *MockUserRepository) GetAllUsers() ([]*models.User, error) {
	return m.GetAllUsersFunc()
}

func TestCreateUserUseCase_Execute(t *testing.T) {
	// Test Case 1: Successful user creation
	t.Run("Success", func(t *testing.T) {
		mockRepo := &MockUserRepository{
			CreateFunc: func(user *models.User) error {
				return nil
			},
		}

		useCase := NewCreateUserUseCase(mockRepo)

		input := CreateUserInputDto{
			Email:     "test@example.com",
			FirstName: "John",
			LastName:  "Doe",
		}

		output, err := useCase.Execute(input)

		assert.NoError(t, err, "Unexpected error")
		assert.NotNil(t, output, "Output should not be nil")
		assert.NotEqual(t, uuid.Nil, output.ID, "ID should not be nil")

		// Additional assertions based on your requirements
		assert.Equal(t, input.Email, output.Email, "Email mismatch")
		assert.Equal(t, input.FirstName, output.FirstName, "FirstName mismatch")
		assert.Equal(t, input.LastName, output.LastName, "LastName mismatch")
	})

	// Test Case 2: User creation failure
	t.Run("Failure", func(t *testing.T) {
		mockRepo := &MockUserRepository{
			CreateFunc: func(user *models.User) error {
				return errors.New("error creating user")
			},
		}

		useCase := NewCreateUserUseCase(mockRepo)

		input := CreateUserInputDto{
			Email:     "test@example.com",
			FirstName: "John",
			LastName:  "Doe",
		}

		output, err := useCase.Execute(input)

		assert.Error(t, err, "Expected error")
		assert.Nil(t, output, "Output should be nil")
	})
}
