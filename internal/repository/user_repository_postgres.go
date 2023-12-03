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
//
// Parameters:
//   - db: A pointer to a *sql.DB representing the PostgreSQL database connection.
//
// Returns:
//   - A pointer to the newly created UserRepositoryPostgres instance.
func NewUserRepositoryPostgres(db *sql.DB) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{DB: db}
}

// Create inserts a new user into the PostgreSQL database.
//
// Parameters:
//   - user: A pointer to a User struct representing the user to be inserted.
//
// Returns:
//   - An error if the operation fails.
func (r *UserRepositoryPostgres) Create(user *models.User) error {
	_, err := r.DB.Exec("INSERT INTO users (id, email, firstName, lastName) VALUES ($1, $2, $3, $4)",
		user.ID, user.Email, user.FirstName, user.LastName)
	if err != nil {
		return err
	}
	return nil
}

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
