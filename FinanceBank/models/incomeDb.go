package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Income struct
type Income struct {
	ID        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Amount    float64   `json:"amount"`
	Category  string    `json:"category"`
	Source    string    `json:"source"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
}

func (i *Income) Create(db *sql.DB) (*Income, error) {
	query := `
	INSERT INTO income (id, amount, category, source, date, user_id, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, NOW()) RETURNING created_at;
	`
	i.ID = uuid.NewString()
	err := db.QueryRow(query, i.ID, i.Amount, i.Category, i.Source, i.Date, i.UserId).Scan(&i.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create income: %w", err)
	}
	return i, nil
}

func GetIncome(db *sql.DB) ([]Income, error) {
	var incomes []Income
	query := `SELECT id, amount, category, source, date, created_at, user_id FROM income`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch incomes: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var i Income
		err := rows.Scan(&i.ID, &i.Amount, &i.Category, &i.Source, &i.Date, &i.CreatedAt, &i.UserId)
		if err != nil {
			return nil, fmt.Errorf("error scanning income row: %w", err)
		}
		incomes = append(incomes, i)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return incomes, nil
}

func GetIncomeById(db *sql.DB, id string) (Income, error) {
	var i Income
	query := `SELECT id, amount, category, source, date, created_at, user_id FROM income WHERE id=$1`

	err := db.QueryRow(query, id).Scan(&i.ID, &i.Amount, &i.Category, &i.Source, &i.Date, &i.CreatedAt, &i.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Income{}, errors.New("income record does not exist")
		}
		return Income{}, fmt.Errorf("error fetching income: %w", err)
	}
	return i, nil
}

func (i *Income) PutValues(db *sql.DB) (int64, error) {
	query := `UPDATE income SET amount=$1, category=$2, source=$3, date=$4 WHERE id=$5`
	result, err := db.Exec(query, i.Amount, i.Category, i.Source, i.Date, i.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to update income: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to check affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return 0, errors.New("no income record found to update")
	}

	return rowsAffected, nil
}

func DeleteIncome(db *sql.DB, id string) error {
	query := `DELETE FROM income WHERE id=$1`
	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete income: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("no income record found to delete")
	}

	return nil
}
