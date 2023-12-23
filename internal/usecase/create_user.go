package usecase

import (
	"github.com/gustavocmaciel/go-mail-api/internal/models"
)

// CreateUserInputDto represents the data structure for creating a user.
type CreateUserInputDto struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// CreateUserOutputDto represents the data structure for the output of a user creation operation.
type CreateUserOutputDto struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
}

// CreateUserUseCase represents a use case for creating a new user.
type CreateUserUseCase struct {
	UserRepository models.UserRepository
}

func NewCreateUserUseCase(userRepository models.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{UserRepository: userRepository}
}

// Execute creates a new user and stores it using the UserRepository
func (u *CreateUserUseCase) Execute(input CreateUserInputDto) (*CreateUserOutputDto, error) {
	user := models.NewUser(input.Email, input.FirstName, input.LastName)

	err := u.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return &CreateUserOutputDto{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}
