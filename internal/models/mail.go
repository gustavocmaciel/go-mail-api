package models

import (
	"time"
)

// Mail represents an email message with the specified attributes.
type Mail struct {
	// User represents the user associated with the email.
	User string `json:"user"`

	// Sender represents the sender's email address.
	Sender string `json:"sender"`

	// Recipients contains a slice of email addresses of the message recipients.
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
