package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// List retrieves users by filter
func (repo Users) List(queryValue string) ([]models.User, error) {
	queryValue = fmt.Sprintf("%%%s%%", queryValue)

	query, err := repo.db.Query(
		"select id, name, nick, email, createdat from users where name like ? or nick like ?",
		queryValue, queryValue,
	)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	var users []models.User

	for query.Next() {
		var user models.User
		if err = query.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
