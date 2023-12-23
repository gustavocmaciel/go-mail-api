package repository

import (
	"database/sql"

	"github.com/gustavocmaciel/go-mail-api/internal/models"
	"github.com/lib/pq"
)

// MailRepositoryPostgres represents a PostgreSQL implementation of the MailRepository interface.
type MailRepositoryPostgres struct {
	DB *sql.DB
}

// NewMailRepositoryPostgres creates a new MailRepositoryPostgres instance.
//
// Parameters:
//   - db: A pointer to a *sql.DB representing the PostgreSQL database connection.
//
// Returns:
//   - A pointer to the newly created MailRepositoryPostgres instance.
func NewMailRepositoryPostgres(db *sql.DB) *MailRepositoryPostgres {
	return &MailRepositoryPostgres{DB: db}
}

// Create inserts a new mail into the PostgreSQL database.
//
// Parameters:
//   - mail: A pointer to a Mail struct representing the mail to be inserted.
//
// Returns:
//   - An error if the operation fails.
func (r *MailRepositoryPostgres) Create(mail *models.Mail) error {
	recipientArray := pq.Array(mail.Recipients) // Use pq.Array here
	_, err := r.DB.Exec(`
		INSERT INTO emails (id, sender, recipients, subject, body, timestamp, email_read, archived)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		mail.ID, mail.Sender, recipientArray,
		mail.Subject, mail.Body, mail.Timestamp, mail.Read, mail.Archived)
	if err != nil {
		return err
	}
	return nil
}

func (r *MailRepositoryPostgres) GetMail(mailID string) (*models.Mail, error) {
	var mail models.Mail
	var recipients pq.StringArray // Use pq.StringArray for PostgreSQL array of strings

	err := r.DB.QueryRow(`
		SELECT id, sender, recipients, subject, body, timestamp, email_read, archived
		FROM emails
		WHERE id = $1;
	`, mailID).Scan(
		&mail.ID, &mail.Sender, &recipients,
		&mail.Subject, &mail.Body, &mail.Timestamp, &mail.Read, &mail.Archived,
	)

	if err != nil {
		return nil, err
	}

	// Convert pq.StringArray to a regular slice of strings
	mail.Recipients = []string(recipients)

	return &mail, nil
}

func (r *MailRepositoryPostgres) Mailbox(user, mailboxName string) ([]*models.Mail, error) {
	var query string
	if mailboxName == "inbox" {
		query = `
			SELECT id, sender, recipients, subject, body, timestamp, email_read, archived
			FROM emails
			WHERE $1 = ANY(recipients);
		`
	} else if mailboxName == "sent" {
		query = `
			SELECT id, sender, recipients, subject, body, timestamp, email_read, archived
			FROM emails
			WHERE sender = $1;
		`
	} else if mailboxName == "archived" {
		query = `
			SELECT id, sender, recipients, subject, body, timestamp, email_read, archived
			FROM emails
			WHERE $1 = ANY(recipients) AND archived = true;
		`
	}

	rows, err := r.DB.Query(query, user)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var mails []*models.Mail
	for rows.Next() {
		var mail models.Mail
		var recipients pq.StringArray

		err := rows.Scan(
			&mail.ID, &mail.Sender, &recipients,
			&mail.Subject, &mail.Body, &mail.Timestamp, &mail.Read, &mail.Archived,
		)
		if err != nil {
			return nil, err
		}

		// Convert pq.StringArray to a regular slice of strings
		mail.Recipients = []string(recipients)

		mails = append(mails, &mail)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return mails, nil
}
