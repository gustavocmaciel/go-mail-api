package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gustavocmaciel/go-mail-api/internal/usecase"
)

type MailHandlers struct {
	CreateMailUseCase *usecase.CreateMailUseCase
	GetMailUseCase    *usecase.GetMailUseCase
	MailboxUserCase   *usecase.MailboxUseCase
}

func NewMailHandlers(createMailUseCase *usecase.CreateMailUseCase, getMailUseCase *usecase.GetMailUseCase, mailboxUseCase *usecase.MailboxUseCase) *MailHandlers {
	return &MailHandlers{
		CreateMailUseCase: createMailUseCase,
		GetMailUseCase:    getMailUseCase,
		MailboxUserCase:   mailboxUseCase,
	}
}

func (u *MailHandlers) CreateMailHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling CreateMail request...")

	var input usecase.CreateMailInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Printf("Error decoding JSON in CreateMailHandler: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Invalid JSON data"))
		return
	}

	log.Printf("Received CreateMail request with input: %+v", input)

	output, err := u.CreateMailUseCase.Execute(input)
	if err != nil {
		log.Printf("Error executing CreateMailUseCase: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("CreateMail request successful.")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (u *MailHandlers) GetMailHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling GetMail request...")

	vars := mux.Vars(r)
	mailID, ok := vars["mail_id"]
	if !ok {
		log.Println("Missing 'mail_id' parameter in GetMail request")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing 'mail_id' parameter"))
		return
	}

	log.Printf("Received GetMail request for MailID: %s", mailID)

	input := usecase.GetMailInputDto{
		MailID: mailID,
	}

	output, err := u.GetMailUseCase.Execute(input)
	if err != nil {
		log.Printf("Error executing GetMailUseCase: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("GetMail request successful.")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (u *MailHandlers) MailboxHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling Mailbox request...")

	vars := mux.Vars(r)
	user, ok := vars["user"]
	if !ok {
		log.Println("Missing 'user' parameter in Mailbox request")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing 'user' parameter"))
		return
	}
	mailbox, ok := vars["mailbox"]
	if !ok {
		log.Println("Missing 'mailbox' parameter in Mailbox request")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing 'mailbox' parameter"))
		return
	}

	log.Printf("Received Mailbox request for User: %s, Mailbox: %s", user, mailbox)

	input := usecase.MailboxInputDto{
		User:        user,
		MailboxName: mailbox,
	}

	output, err := u.MailboxUserCase.Execute(input)
	if err != nil {
		log.Printf("Error executing MailboxUserCase: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Mailbox request successful.")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
