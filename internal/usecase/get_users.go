package usecase

import (
	"github.com/gustavocmaciel/go-mail-api/internal/models"
)

type GetUsersOutputDto struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
}

type GetUsersUseCase struct {
	UserRepository models.UserRepository
}

func NewGetUsersUseCase(userRepository models.UserRepository) *GetUsersUseCase {
	return &GetUsersUseCase{UserRepository: userRepository}
}

func (u *GetUsersUseCase) Execute() ([]*GetUsersOutputDto, error) {
	users, err := u.UserRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var usersOutput []*GetUsersOutputDto
	for _, user := range users {
		usersOutput = append(usersOutput, &GetUsersOutputDto{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		})
	}
	return usersOutput, nil
}
