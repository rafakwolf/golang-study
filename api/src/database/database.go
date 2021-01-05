package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySql driver
)

// Connect opens database connection
func Connect() (*sql.DB, error) {
	db, error := sql.Open("mysql", config.ConnString)
	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		db.Close()
		return nil, error
	}

	return db, nil
}
