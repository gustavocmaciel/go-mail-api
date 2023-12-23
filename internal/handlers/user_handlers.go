package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gustavocmaciel/go-mail-api/internal/usecase"
)

type UserHandlers struct {
	CreateUserUseCase *usecase.CreateUserUseCase
	GetUsersUseCase   *usecase.GetUsersUseCase
}

func NewUserHandlers(createUserUseCase *usecase.CreateUserUseCase, getUsersUseCase *usecase.GetUsersUseCase) *UserHandlers {
	return &UserHandlers{
		CreateUserUseCase: createUserUseCase,
		GetUsersUseCase:   getUsersUseCase,
	}
}

func (u *UserHandlers) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling CreateUser request...")

	var input usecase.CreateUserInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Printf("Error decoding JSON in CreateUserHandler: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Invalid JSON data"))
		return
	}

	log.Printf("Received CreateUser request: %+v", input)

	output, err := u.CreateUserUseCase.Execute(input)
	if err != nil {
		log.Printf("Error executing CreateUserUseCase: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("CreateUser request processed successfully.")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (u *UserHandlers) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling GetUsers request...")

	output, err := u.GetUsersUseCase.Execute()
	if err != nil {
		log.Printf("Error executing GetUsersUseCase: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("GetUsers request processed successfully.")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
