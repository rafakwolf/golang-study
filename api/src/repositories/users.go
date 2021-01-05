package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users is used to communicate with database
type Users struct {
	db *sql.DB
}

// NewRepository initialize the Users repository
func NewRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create adds a new user to the db and returns the new ID
func (repo Users) Create(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare(
		"insert into users(name, nick, email, password) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	res, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}
