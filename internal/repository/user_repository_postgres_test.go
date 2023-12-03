package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/gustavocmaciel/go-mail-api/internal/models"
)

func TestCreateUser(t *testing.T) {
	// Create a new mock database and UserRepositoryPostgres instance
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()
	repo := NewUserRepositoryPostgres(db)
	// Test Case 1: Successful insertion
	user := &models.User{
		ID:        uuid.New(),
		Email:     "test@example.com",
		FirstName: "John",
		LastName:  "Doe",
	}
	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.ID, user.Email, user.FirstName, user.LastName).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err = repo.Create(user)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Test Case 2: Database error during insertion
	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.ID, user.Email, user.FirstName, user.LastName).
		WillReturnError(errors.New("database error"))
	err = repo.Create(user)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestGetAllUsers(t *testing.T) {
	// Create a new mock database and UserRepositoryPostgres instance
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()
	repo := NewUserRepositoryPostgres(db)
	// Test Case 1: Successful retrieval
	rows := sqlmock.NewRows([]string{"id", "email", "first_name", "last_name"}).
		AddRow(uuid.New(), "test1@example.com", "John", "Doe").
		AddRow(uuid.New(), "test2@example.com", "Jane", "Doe")
	mock.ExpectQuery("SELECT id, email, first_name, last_name FROM users").
		WillReturnRows(rows)
	users, err := repo.GetAllUsers()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(users) != 2 {
		t.Errorf("Expected 2 users, but got %d", len(users))
	}
	// Test Case 2: Database error during retrieval
	mock.ExpectQuery("SELECT id, email, first_name, last_name FROM users").
		WillReturnError(errors.New("database error"))
	_, err = repo.GetAllUsers()
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
