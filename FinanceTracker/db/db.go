package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", "user=postgres dbname=test password=admin sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("Could not connect to the database: %v", err))
	}
	err = createTable()
	if err != nil {
		panic(fmt.Sprintf("Error creating table: %v", err))
	}
	println("New table Created")
}

func createTable() error {
	query := `CREATE TABLE IF NOT EXISTS income (
    id SERIAL PRIMARY KEY,
    amount DECIMAL(10,2) NOT NULL,
    category VARCHAR(50) NOT NULL,
    source VARCHAR(100),
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
)`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	query = `CREATE TABLE IF NOT EXISTS expense (
    id SERIAL PRIMARY KEY,
    amount DECIMAL(10,2) NOT NULL,
    category VARCHAR(50) NOT NULL,
    source VARCHAR(100),
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
)`
	stmt, err = DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}
