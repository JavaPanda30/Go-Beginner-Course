package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Expense struct
type Expense struct {
	ID        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Amount    float64   `json:"amount"`
	Category  string    `json:"category"`
	Source    string    `json:"source"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
}

func (e *Expense) Create(db *sql.DB) (*Expense, error) {
	query := `
	INSERT INTO expense (id, amount, category, source, date, user_id, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7);
	`
	e.ID = uuid.NewString()
	e.CreatedAt = time.Now()

	_, err := db.Exec(query, e.ID, e.Amount, e.Category, e.Source, e.Date, e.UserId, e.CreatedAt)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func GetExp(db *sql.DB) ([]Expense, error) {
	var expenses []Expense
	query := `SELECT id, amount, category, source, date, created_at, user_id FROM expense`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e Expense
		err := rows.Scan(&e.ID, &e.Amount, &e.Category, &e.Source, &e.Date, &e.CreatedAt, &e.UserId)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, e)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return expenses, nil
}

func GetIdExp(db *sql.DB, id string) (Expense, error) {
	var e Expense
	query := `SELECT id, amount, category, source, date, created_at, user_id FROM expense WHERE id=$1`
	err := db.QueryRow(query, id).Scan(&e.ID, &e.Amount, &e.Category, &e.Source, &e.Date, &e.CreatedAt, &e.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return Expense{}, errors.New("expense record does not exist")
		}
		return Expense{}, err
	}
	return e, nil
}

func (e *Expense) PutValues(db *sql.DB) (int64, error) {
	query := `UPDATE expense SET amount=$1, category=$2, source=$3, date=$4 WHERE id=$5`
	result, err := db.Exec(query, e.Amount, e.Category, e.Source, e.Date, e.ID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func DeleteExp(db *sql.DB, id string) error {
	query := `DELETE FROM expense WHERE id=$1`
	_, err := db.Exec(query, id)
	return err
}
