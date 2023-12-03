package mocks

import (
	"time"

	"github.com/gustavocmaciel/go-mail-api/internal/models"
)

var Mails = []models.Mail{
	{
		User:       Users[0].Email,
		Sender:     Users[0].Email,
		Recipients: []string{Users[1].Email},
		Subject:    "Hello Jane",
		Body:       "Hi Jane, how are you?",
		Timestamp:  time.Now(),
		Read:       false,
		Archived:   false,
	},
	{
		User:       Users[1].Email,
		Sender:     Users[1].Email,
		Recipients: []string{Users[0].Email},
		Subject:    "Re: Hello Jane",
		Body:       "Hi John, I'm good. How about you?",
		Timestamp:  time.Now(),
		Read:       false,
		Archived:   false,
	},
	{
		User:       Users[2].Email,
		Sender:     Users[2].Email,
		Recipients: []string{Users[0].Email},
		Subject:    "Greetings John",
		Body:       "Dear John, I hope you are doing well.",
		Timestamp:  time.Now(),
		Read:       false,
		Archived:   false,
	},
	{
		User:       Users[0].Email,
		Sender:     Users[0].Email,
		Recipients: []string{Users[2].Email},
		Subject:    "Re: Greetings John",
		Body:       "Hi Michael, thanks for the wishes.",
		Timestamp:  time.Now(),
		Read:       false,
		Archived:   false,
	},
	{
		User:       Users[2].Email,
		Sender:     Users[2].Email,
		Recipients: []string{Users[1].Email},
		Subject:    "Hello Jane",
		Body:       "Hi Jane, how have you been?",
		Timestamp:  time.Now(),
		Read:       false,
		Archived:   false,
	},
}
