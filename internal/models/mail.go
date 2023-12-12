package models

import (
	"time"

	"github.com/google/uuid"
)

// MailRepository defines the methods that a repository for managing mail data should implement.
type MailRepository interface {
	// Create inserts a new mail into the repository.
	Create(mail *Mail) error
	// GetMail retrieves a specific mail from the repository based on its ID.
	GetMail(mailID string) (*Mail, error)
	// Mailbox retrieves a list of mails from a specific mailbox for a given user.
	Mailbox(user, mailboxName string) ([]*Mail, error)
}

// Mail represents an email message with the specified attributes.
type Mail struct {
	// ID is the unique identifier for the email.
	ID string `json:"id"`
	// Sender represents the sender's email address.
	Sender string `json:"sender"`
	// Recipients contains a slice of email addresses of receivers.
	Recipients []string `json:"recipients"`
	// Subject represents the subject of the email.
	Subject string `json:"subject"`
	// Body contains the content of the email.
	Body string `json:"body"`
	// Timestamp represents the time the email was sent.
	Timestamp time.Time `json:"timestamp"`
	// Read specifies whether the email has been read or not.
	Read bool `json:"read"`
	// Archived specifies whether the email has been archived or not.
	Archived bool `json:"archived"`
}

func NewMail(sender string, recipients []string, subject, body string) *Mail {
	return &Mail{
		ID:         uuid.New().String(),
		Sender:     sender,
		Recipients: recipients,
		Subject:    subject,
		Body:       body,
		Timestamp:  time.Now(),
		Read:       false,
		Archived:   false,
	}
}
