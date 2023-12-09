package models

import (
	"time"

	"github.com/google/uuid"
)

type MailRepository interface {
	// Inserts a new mail into the repository.
	// Returns an error if the operation fails.
	Create(mail *Mail) error
	// GetMail retrieves a specific mail from the repository based on its ID and user.
	// Parameters:
	//   - mailID: The ID of the mail to retrieve.
	// Returns the requested mail and an error if the operation fails.
	GetMail(mailID uuid.UUID) (*Mail, error)
	// Retrieves a list of mails from a specific mailbox for a given user.
	// Parameters:
	//   - user: The email of the user who owns the mailbox.
	//   - mailboxName: The name of the mailbox (e.g., Inbox, Sent, Archived).
	// Returns a slice of Mail and an error if the operation fails.
	Mailbox(user, mailboxName string) ([]*Mail, error)
}

// Mail represents an email message with the specified attributes.
type Mail struct {
	// ID is the unique identifier for the email.
	ID uuid.UUID `json:"id"`
	// Sender represents the sender's email address.
	Sender string `json:"sender"`
	// Recipients contains a slice of email addresses associated with the email.
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

// Creates a new Mail instance
func NewMail(sender string, recipients []string, subject, body string, timestamp time.Time, read, archived bool) *Mail {
	return &Mail{
		ID:         uuid.New(),
		Sender:     sender,
		Recipients: recipients,
		Subject:    subject,
		Body:       body,
		Timestamp:  timestamp,
		Read:       false,
		Archived:   false,
	}

}
