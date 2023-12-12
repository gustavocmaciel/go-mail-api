package usecase

import (
	"time"

	"github.com/gustavocmaciel/go-mail-api/internal/models"
)

type GetMailInputDto struct {
	MailID string `json:"mail_id"`
}

type GetMailOutputDto struct {
	ID         string    `json:"id"`
	Sender     string    `json:"sender"`
	Recipients []string  `json:"recipients"`
	Subject    string    `json:"subject"`
	Body       string    `json:"body"`
	Timestamp  time.Time `json:"timestamp"`
	Read       bool      `json:"read"`
	Archived   bool      `json:"archived"`
}

type GetMailUseCase struct {
	MailRepository models.MailRepository
}

func NewGetMailUseCase(mailRepository models.MailRepository) *GetMailUseCase {
	return &GetMailUseCase{MailRepository: mailRepository}
}

func (u *GetMailUseCase) Execute(input GetMailInputDto) (*GetMailOutputDto, error) {
	mail, err := u.MailRepository.GetMail(input.MailID)
	if err != nil {
		return nil, err
	}
	return &GetMailOutputDto{
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
