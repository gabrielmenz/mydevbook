package repositories

import (
	"database/sql"
	"goapi/src/models"
)

// this represents a users repository
type users struct {
	df *sql.DB
}

// This func creates a new users repository
func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

// This adds user to db
func (repository users) Creates(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err

	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	lastIDInserted, err := result.lastIDInserted()
	if err != nil {
		return 0, err
	}

	return uint64(lastIDInserted), nil
}
