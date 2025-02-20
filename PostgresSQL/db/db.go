package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	var err error

	connStr := "host=localhost port=5432 user=postgres password=admin dbname=test sslmode=disable"

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("could not connect to database: %v", err)
	}
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("could not ping database: %v", err)
	}

	fmt.Println("DB CONNECTED")

	if err := createTables(); err != nil {
		return fmt.Errorf("could not create tables: %v", err)
	}
	fmt.Println("Tables Created Successfully")
	return nil
}

func createTables() error {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
		);`
	_, err := DB.Exec(createUserTable)
	if err != nil {
		return fmt.Errorf("error creating users table: %v", err)
	}
	fmt.Println("Users table created successfully or already exists.")

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			datetime TIMESTAMP NOT NULL,
			user_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(id)
			);`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		return fmt.Errorf("error creating events table: %v", err)
	}
	fmt.Println("Events table created successfully or already exists.")
	registrationTable := `
			CREATE TABLE IF NOT EXISTS registration (
			id SERIAL PRIMARY KEY,
			event_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);
		`
	_, err = DB.Exec(registrationTable)
	if err != nil {
		return fmt.Errorf("error creating registration table: %v", err)
	}
	return nil
}
