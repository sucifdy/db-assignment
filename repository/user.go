package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
	FetchByID(id int) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) Add(user model.User) error {
	// Insert a new user into the users table
	_, err := u.db.Exec(
		"INSERT INTO users (username, password) VALUES ($1, $2)",
		user.Username, user.Password,
	)
	return err
}

func (u *userRepository) CheckAvail(user model.User) error {
	// Check if a user with the given username and password exists
	row := u.db.QueryRow("SELECT id FROM users WHERE username = $1 AND password = $2", user.Username, user.Password)

	var id int
	err := row.Scan(&id)
	if err != nil {
		// If no rows are found, return an error indicating user is not available
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found")
		}
		// Return any other errors
		return err
	}

	return nil
}

func (u *userRepository) FetchByID(id int) (*model.User, error) {
	row := u.db.QueryRow("SELECT id, username, password FROM users WHERE id = $1", id)

	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
