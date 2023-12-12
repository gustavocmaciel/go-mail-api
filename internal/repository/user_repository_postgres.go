package repository

import (
	"database/sql"

	"github.com/gustavocmaciel/go-mail-api/internal/models"
)

// UserRepositoryPostgres represents a PostgreSQL implementation of the UserRepository interface.
type UserRepositoryPostgres struct {
	DB *sql.DB
}

// NewUserRepositoryPostgres creates a new UserRepositoryPostgres instance.
func NewUserRepositoryPostgres(db *sql.DB) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{DB: db}
}

// Create inserts a new user into the PostgreSQL database.
func (r *UserRepositoryPostgres) Create(user *models.User) error {
	_, err := r.DB.Exec("INSERT INTO users (id, email, first_name, last_name) VALUES ($1, $2, $3, $4)",
		user.ID, user.Email, user.FirstName, user.LastName)
	if err != nil {
		return err
	}
	return nil
}

// GetAllUsers retrieves all users from the  database.
func (r *UserRepositoryPostgres) GetAllUsers() ([]*models.User, error) {
	rows, err := r.DB.Query("SELECT id, email, first_name, last_name FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
