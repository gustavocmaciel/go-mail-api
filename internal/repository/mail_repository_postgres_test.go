package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/gustavocmaciel/go-mail-api/internal/models"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestMailRepositoryPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	repo := NewMailRepositoryPostgres(db)

	// Mock data
	mailID := uuid.New()
	mail := &models.Mail{
		ID:         mailID,
		Sender:     "sender@example.com",
		Recipients: []string{"recipient@example.com"},
		Subject:    "Test Subject",
		Body:       "Test Body",
		Timestamp:  time.Now(),
		Read:       false,
		Archived:   false,
	}
	// Mock expected SQL query
	mock.ExpectExec(`INSERT INTO emails`).WithArgs(
		mail.ID, mail.Sender, pq.Array(mail.Recipients),
		mail.Subject, mail.Body, mail.Timestamp, mail.Read, mail.Archived,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	// Test the Create method
	err = repo.Create(mail)
	assert.NoError(t, err)
}

func TestMailRepositoryPostgres_GetMail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	repo := NewMailRepositoryPostgres(db)

	// Mock data
	mailID := uuid.New()
	mail := &models.Mail{
		ID:         mailID,
		Sender:     "sender@example.com",
		Recipients: []string{"recipient@example.com"},
		Subject:    "Test Subject",
		Body:       "Test Body",
		Timestamp:  time.Now(),
		Read:       false,
		Archived:   false,
	}

	// Convert recipients to pq.StringArray for PostgreSQL array syntax
	pgArrayRecipients := pq.Array(mail.Recipients)

	// Mock expected SQL query and result
	rows := sqlmock.NewRows([]string{"id", "sender", "recipients", "subject", "body", "timestamp", "email_read", "archived"}).
		AddRow(mail.ID, mail.Sender, pgArrayRecipients, mail.Subject, mail.Body, mail.Timestamp, mail.Read, mail.Archived)

	mock.ExpectQuery(`SELECT id, sender, recipients, subject, body, timestamp, email_read, archived`).WithArgs(mailID).WillReturnRows(rows)

	// Test the GetMail method
	retrievedMail, err := repo.GetMail(mailID)
	assert.NoError(t, err)
	assert.Equal(t, mail, retrievedMail)
}

func TestMailRepositoryPostgres_Mailbox(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	repo := NewMailRepositoryPostgres(db)

	// Mock data
	user := "user@example.com"
	mailboxName := "Inbox"
	mails := []*models.Mail{
		{
			ID:         uuid.New(),
			Sender:     "sender1@example.com",
			Recipients: []string{user},
			Subject:    "Test Subject 1",
			Body:       "Test Body 1",
			Timestamp:  time.Now(),
			Read:       false,
			Archived:   false,
		},
		{
			ID:         uuid.New(),
			Sender:     "sender2@example.com",
			Recipients: []string{user},
			Subject:    "Test Subject 2",
			Body:       "Test Body 2",
			Timestamp:  time.Now(),
			Read:       true,
			Archived:   false,
		},
	}

	// Convert recipients to pq.StringArray for PostgreSQL array syntax
	pgArrayRecipients := pq.Array([]string{user})

	// Mock expected SQL query and result
	rows := sqlmock.NewRows([]string{"id", "sender", "recipients", "subject", "body", "timestamp", "email_read", "archived"}).
		AddRow(mails[0].ID, mails[0].Sender, pgArrayRecipients, mails[0].Subject, mails[0].Body, mails[0].Timestamp, mails[0].Read, mails[0].Archived).
		AddRow(mails[1].ID, mails[1].Sender, pgArrayRecipients, mails[1].Subject, mails[1].Body, mails[1].Timestamp, mails[1].Read, mails[1].Archived)

	mock.ExpectQuery(`SELECT id, sender, recipients, subject, body, timestamp, email_read, archived`).WithArgs(pgArrayRecipients, mailboxName).WillReturnRows(rows)

	// Test the Mailbox method
	retrievedMails, err := repo.Mailbox(user, mailboxName)
	assert.NoError(t, err)
	assert.Equal(t, mails, retrievedMails)
}
