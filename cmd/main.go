package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gustavocmaciel/go-mail-api/internal/handlers"
	"github.com/gustavocmaciel/go-mail-api/internal/repository"
	"github.com/gustavocmaciel/go-mail-api/internal/usecase"
)

func main() {
	log.Println("Starting the application...")

	// Establish database connection
	log.Println("Connecting to the database...")
	db, err := sql.Open("postgres", "user=admin password=adminpassword host=host.docker.internal port=5432 dbname=maildatabase sslmode=disable")
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	defer db.Close()
	log.Println("Database connection established successfully.")

	// User setup
	log.Println("Setting up user components...")
	userRepository := repository.NewUserRepositoryPostgres(db)
	createUserUseCase := usecase.NewCreateUserUseCase(userRepository)
	getUsersUseCase := usecase.NewGetUsersUseCase(userRepository)
	userHandlers := handlers.NewUserHandlers(createUserUseCase, getUsersUseCase)
	log.Println("User components set up successfully.")

	// Mail setup
	log.Println("Setting up mail components...")
	mailRepository := repository.NewMailRepositoryPostgres(db)
	createMailUseCase := usecase.NewCreateMailUseCase(mailRepository)
	getMailUseCase := usecase.NewGetMailUseCase(mailRepository)
	mailboxUseCase := usecase.NewMailboxUseCase(mailRepository)
	mailHandlers := handlers.NewMailHandlers(createMailUseCase, getMailUseCase, mailboxUseCase)
	log.Println("Mail components set up successfully.")

	// HTTP server setup
	log.Println("Setting up HTTP server and routes...")
	router := mux.NewRouter()
	// Routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Mailbox")
	})
	router.HandleFunc("/create_user", userHandlers.CreateUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/get_users", userHandlers.GetUsersHandler).Methods(http.MethodGet)
	router.HandleFunc("/create_mail", mailHandlers.CreateMailHandler).Methods(http.MethodPost)
	router.HandleFunc("/get_mail/{mail_id}", mailHandlers.GetMailHandler).Methods(http.MethodGet)
	router.HandleFunc("/mailbox/{user}/{mailbox}", mailHandlers.MailboxHandler).Methods(http.MethodGet)
	log.Println("HTTP server and routes set up successfully.")

	// Start the HTTP server
	log.Println("Starting the API server on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting the API server: %v", err)
	}
}
