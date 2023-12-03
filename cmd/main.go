package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gustavocmaciel/go-mail-api/internal/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Mailbox")
	})
	router.HandleFunc("/users", handlers.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/mails", handlers.GetAllMails).Methods(http.MethodGet)
	router.HandleFunc("/mail/{user}/{mailbox}", handlers.GetMail).Methods(http.MethodGet)
	router.HandleFunc("/add_user", handlers.AddUser).Methods(http.MethodPost)
	log.Println("API is running!")
	http.ListenAndServe("localhost:8080", router)
	// http.ListenAndServe(":8080", router) // for production environment
}
