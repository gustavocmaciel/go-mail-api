package models

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Mock implementation of MailRepository for testing purposes.
type MockMailRepository struct {
	Emails map[uuid.UUID]*Mail
}

func (m *MockMailRepository) Create(mail *Mail) error {
	m.Emails[mail.ID] = mail
	return nil
}

func (m *MockMailRepository) GetMail(mailID uuid.UUID) (*Mail, error) {
	mail, ok := m.Emails[mailID]
	if !ok {
		return nil, nil
	}
	return mail, nil
}

func (m *MockMailRepository) Mailbox(user string, mailboxName string) ([]*Mail, error) {
	var result []*Mail
	for _, mail := range m.Emails {
		if mail.Sender == user || contains(mail.Recipients, user) {
			result = append(result, mail)
		}
	}
	return result, nil
}

// Helper function to check if a string is in a slice of strings.
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// Test cases for the MailRepository interface
func TestMailRepository(t *testing.T) {
	// Create a sample mail
	sampleMail := NewMail("sender@example.com", []string{"recipient@example.com"}, "Test Subject", "Test Body", time.Now(), false, false)
	// Create a mock repository
	mockRepo := &MockMailRepository{
		Emails: make(map[uuid.UUID]*Mail),
	}

	// Test Case 1: Create a mail
	err := mockRepo.Create(sampleMail)
	assert.NoError(t, err)

	// Test Case 2: Get a mail by ID
	retrievedMail, err := mockRepo.GetMail(sampleMail.ID)
	assert.NoError(t, err)
	assert.Equal(t, sampleMail, retrievedMail)

	// Test Case 3: Get a non-existent mail by ID
	nonExistentMailID := uuid.New()
	nonExistentMail, err := mockRepo.GetMail(nonExistentMailID)
	assert.NoError(t, err)
	assert.Nil(t, nonExistentMail)

	// Test Case 4: Retrieve mails from a mailbox
	mails, err := mockRepo.Mailbox("sender@example.com", "Inbox")
	assert.NoError(t, err)
	assert.Len(t, mails, 1)
	assert.Equal(t, sampleMail, mails[0])
}
