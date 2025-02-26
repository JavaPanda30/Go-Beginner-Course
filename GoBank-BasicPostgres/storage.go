package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int64) error
	UpdateAccount(*Account) error
	GetAccountByID(int64) (*Account, error)
}

type PostgresStore struct {
	DB *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=admin sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{DB: db}, nil
}

func (s *PostgresStore) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		acc_num SERIAL UNIQUE,
		balance FLOAT DEFAULT 0.0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := s.DB.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `INSERT INTO account (first_name, last_name, balance) VALUES ($1, $2, $3) RETURNING id, acc_num`
	return s.DB.QueryRow(query, acc.FirstName, acc.LastName, acc.Balance).Scan(&acc.ID, &acc.Number)
}

func (s *PostgresStore) DeleteAccount(id int64) error {
	query := `DELETE FROM account WHERE id = $1`
	_, err := s.DB.Exec(query, id)
	return err
}

func (s *PostgresStore) UpdateAccount(acc *Account) error {
	query := `UPDATE account SET balance = $1 WHERE id = $2`
	_, err := s.DB.Exec(query, acc.Balance, acc.ID)
	return err
}

func (s *PostgresStore) GetAccountByID(id int64) (*Account, error) {
	query := `SELECT id, first_name, last_name, acc_num, balance FROM account WHERE id = $1`
	acc := &Account{}
	if err := s.DB.QueryRow(query, id).Scan(&acc.ID, &acc.FirstName, &acc.LastName, &acc.Number, &acc.Balance); err != nil {
		return nil, err
	}
	return acc, nil
}
