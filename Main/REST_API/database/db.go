package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// OpenDatabase opens a connection to the PostgreSQL database using the provided connection string,It returns an error if the connection could not be established.
func OpenDatabase() error {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=admin dbname=person sslmode=disable")
	if err != nil {
		return err
	}
	return nil
}

// CloseDatabase closes the database connection if it is open,It returns an error if the close operation fails.

func CloseDatabase() error {
	if DB == nil {
		return nil
	}
	return DB.Close()
}
