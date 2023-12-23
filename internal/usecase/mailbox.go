package usecase

import (
	"time"

	"github.com/gustavocmaciel/go-mail-api/internal/models"
)

type MailboxInputDto struct {
	User        string `json:"user"`
	MailboxName string `json:"mailbox"`
}

type MailboxOutputDto struct {
	ID         string    `json:"id"`
	Sender     string    `json:"sender"`
	Recipients []string  `json:"recipients"`
	Subject    string    `json:"subject"`
	Body       string    `json:"body"`
	Timestamp  time.Time `json:"timestamp"`
	Read       bool      `json:"read"`
	Archived   bool      `json:"archived"`
}

type MailboxUseCase struct {
	MailRepository models.MailRepository
}

func NewMailboxUseCase(mailRepository models.MailRepository) *MailboxUseCase {
	return &MailboxUseCase{MailRepository: mailRepository}
}

func (u *MailboxUseCase) Execute(input MailboxInputDto) ([]*MailboxOutputDto, error) {
	mails, err := u.MailRepository.Mailbox(input.User, input.MailboxName)
	if err != nil {
		return nil, err
	}

	var mailboxOutput []*MailboxOutputDto
	for _, mail := range mails {
		mailboxOutput = append(mailboxOutput, &MailboxOutputDto{
			ID:         mail.ID,
			Sender:     mail.Sender,
			Recipients: mail.Recipients,
			Subject:    mail.Subject,
			Body:       mail.Body,
			Timestamp:  mail.Timestamp,
			Read:       mail.Read,
			Archived:   mail.Archived,
		})
	}
	return mailboxOutput, nil
}
