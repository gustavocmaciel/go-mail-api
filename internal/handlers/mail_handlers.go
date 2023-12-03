package handlers

import (
	"encoding/json"
	"net/http"
	"slices"

	"github.com/gorilla/mux"
	"github.com/gustavocmaciel/go-mail-api/internal/mocks"
)

func GetAllMails(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Mails)
}

func Mailbox(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Mails)
}

func GetMail(w http.ResponseWriter, r *http.Request) {
	// Read dynamic user parameter and mailbox parameter
	vars := mux.Vars(r)
	user := vars["user"]
	mailbox := vars["mailbox"]

	if mailbox == "inbox" {
		for _, mail := range mocks.Mails {
			if slices.Contains(mail.Recipients, user) {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(mail)
				break
			}
		}
	} else if mailbox == "sent" {
		for _, mail := range mocks.Mails {
			if mail.Sender == user {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(mail)
				break
			}
		}
	} else if mailbox == "archive" {
		for _, mail := range mocks.Mails {
			if slices.Contains(mail.Recipients, user) && mail.Archived {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(mail)
				break
			}
		}
	}
}
