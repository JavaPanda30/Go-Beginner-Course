package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "user=aman password=admin port=1234 sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	if err := createTables(); err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}

	log.Println("Database connected and tables initialized successfully")
}

func createTables() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			account_id UUID,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password_hash TEXT NOT NULL,
			currency VARCHAR(10) NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		);`,
		`CREATE TABLE IF NOT EXISTS income (
			id UUID PRIMARY KEY,
			amount DECIMAL(10,2) NOT NULL,
			category VARCHAR(50) NOT NULL,
			source VARCHAR(100),
			date DATE NOT NULL,
			created_at TIMESTAMP DEFAULT NOW(),
			user_id UUID NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS expense (
			id UUID PRIMARY KEY,
			amount DECIMAL(10,2) NOT NULL,
			category VARCHAR(50) NOT NULL,
			source VARCHAR(100),
			date DATE NOT NULL,
			created_at TIMESTAMP DEFAULT NOW(),
			user_id UUID NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS transfers (
			id UUID PRIMARY KEY,
			from_user UUID NOT NULL,
			to_user UUID NOT NULL,
			amount DECIMAL(10,2) NOT NULL,
			time TIMESTAMP DEFAULT NOW(),
			FOREIGN KEY(from_user) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY(to_user) REFERENCES users(id) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS accounts (
			id UUID PRIMARY KEY,
			user_id UUID NOT NULL,
			amount DECIMAL(10,2) NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,
	}

	for _, query := range queries {
		if _, err := DB.Exec(query); err != nil {
			return fmt.Errorf("error executing query: %v", err)
		}
	}
	return nil
}
