package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		panic("Could not connect to database: " + err.Error())
	}
	fmt.Println("DB CONNECTED")
	if err := createTables(); err != nil {
		panic("Could not create tables: " + err.Error())
	}
}

func createTables() error {
	registrationTable := `
	CREATE TABLE IF NOT EXISTS registration(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id INTEGER,
	user_id INTEGER,
	FOREIGN KEY(event_id) REFERENCES events(id)
	FOREIGN KEY(user_id) REFERENCES user(id))
	`
	_, err := DB.Exec(registrationTable)
	if err != nil {
		return err
	}
	createUserTabe := `
	CREATE TABLE IF NOT EXISTS user (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL 
	)`
	_, err = DB.Exec(createUserTabe)
	if err != nil {
		return err
	}
	fmt.Println("User table created successfully or already exists.")
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES user(id)
	);`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		return err
	}
	fmt.Println("Events table created successfully or already exists.")
	return nil
}
