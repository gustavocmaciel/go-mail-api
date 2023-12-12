package usecase

import (
	"time"

	"github.com/gustavocmaciel/go-mail-api/internal/models"
)

// CreateMailInputDto represents the data structure for creating a new mail.
type CreateMailInputDto struct {
	Sender     string   `json:"sender"`
	Recipients []string `json:"recipients"`
	Subject    string   `json:"subject"`
	Body       string   `json:"body"`
}

// CreateMailOutputDto represents the data structure for the output of creating a new mail.
type CreateMailOutputDto struct {
	ID         string    `json:"id"`
	Sender     string    `json:"sender"`
	Recipients []string  `json:"recipients"`
	Subject    string    `json:"subject"`
	Body       string    `json:"body"`
	Timestamp  time.Time `json:"timestamp"`
	Read       bool      `json:"read"`
	Archived   bool      `json:"archived"`
}

// CreateMailUseCase represents a use case for creating a new mail.
type CreateMailUseCase struct {
	MailRepository models.MailRepository
}

func NewCreateMailUseCase(mailRepository models.MailRepository) *CreateMailUseCase {
	return &CreateMailUseCase{MailRepository: mailRepository}
}

// Execute creates a new mail and stores it using the MailRepository.
func (u *CreateMailUseCase) Execute(input CreateMailInputDto) (*CreateMailOutputDto, error) {
	mail := models.NewMail(input.Sender, input.Recipients, input.Subject, input.Body)

	err := u.MailRepository.Create(mail)
	if err != nil {
		return nil, err
	}
	return &CreateMailOutputDto{
		ID:         mail.ID,
		Sender:     mail.Sender,
		Recipients: mail.Recipients,
		Subject:    mail.Subject,
		Body:       mail.Body,
		Timestamp:  mail.Timestamp,
		Read:       mail.Read,
		Archived:   mail.Archived,
	}, nil
}
