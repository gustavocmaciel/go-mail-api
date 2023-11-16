package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gustavocmaciel/go-mail-api/pkg/mocks"
	"github.com/gustavocmaciel/go-mail-api/pkg/models"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Users)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	// Read request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error reading request body"))
		return
	}

	// Unmarshal JSON data
	var newUser models.User
	if err := json.Unmarshal(body, &newUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON data"))
		return
	}

	// Add the new user to the users slice
	mocks.Users = append(mocks.Users, newUser)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User added successfully"))
}
