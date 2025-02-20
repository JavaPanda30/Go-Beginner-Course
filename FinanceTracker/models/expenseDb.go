package models

import (
	"database/sql"
	"errors"
	"time"
)

// Expense struct
type Expense struct {
	ID        int64     `json:"id"`
	Amount    float64   `json:"amount"`
	Category  string    `json:"category"`
	Source    string    `json:"source"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
}

func (i *Expense) Create(db *sql.DB) (*Expense, error) {
	query := `
	INSERT INTO expense (amount, category, source, date) 
	VALUES ($1, $2, $3, $4) 
	RETURNING id, created_at;
	`
	err := db.QueryRow(query, i.Amount, i.Category, i.Source, i.Date).
		Scan(&i.ID, &i.CreatedAt)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func GetExp(db *sql.DB) ([]Expense, error) {
	var expense []Expense
	query := `
	SELECT * FROM expense
	`
	rows, err := db.Query(query)
	if err != nil {
		return []Expense{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var i Expense
		err := rows.Scan(&i.ID, &i.Amount, &i.Category, &i.Source, &i.Date, &i.CreatedAt)
		if err != nil {
			return []Expense{}, err
		}
		expense = append(expense, i)
	}
	if err = rows.Err(); err != nil {
		return []Expense{}, err
	}
	return expense, nil
}

func GetIdExp(db *sql.DB, id int64) (Expense, error) {
	expense, err := GetExp(db)
	if err != nil {
		return Expense{}, err
	}
	for _, i := range expense {
		if i.ID == id {
			return i, nil
		}
	}
	return Expense{}, errors.New("user does not Exist")
}

func (i *Expense) PutValues(db *sql.DB) (int64, error) {
	query := `UPDATE expense SET amount=$1, category=$2, source=$3, date=$4 WHERE id=$5`
	result, err := db.Exec(query, i.Amount, i.Category, i.Source, i.Date, i.ID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func DeleteExp(db *sql.DB, id int64) error {
	query := `DELETE FROM expense where id=$1`
	_, err := db.Exec(query, id)
	return err
}
