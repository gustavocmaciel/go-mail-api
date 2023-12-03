package mocks

import "github.com/gustavocmaciel/go-mail-api/internal/models"

var Users = []models.User{
	{
		Email:     "johndoe@example.com",
		FirstName: "John",
		LastName:  "Doe",
	},
	{
		Email:     "janesmith@example.com",
		FirstName: "Jane",
		LastName:  "Smith",
	},
	{
		Email:     "michaeljohnson@example.com",
		FirstName: "Michael",
		LastName:  "Johnson",
	},
}
