package models

import (
	"database/sql"
	"errors"
	"time"
)

// Income struct
type Income struct {
	ID        int64     `json:"id"`
	Amount    float64   `json:"amount"`
	Category  string    `json:"category"`
	Source    string    `json:"source"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
}

func (i *Income) Create(db *sql.DB) (*Income, error) {
	query := `
	INSERT INTO income (amount, category, source, date) 
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

func Get(db *sql.DB) ([]Income, error) {
	var income []Income
	query := `
	SELECT * FROM income
	`
	rows, err := db.Query(query)
	if err != nil {
		return []Income{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var i Income
		err := rows.Scan(&i.ID, &i.Amount, &i.Category, &i.Source, &i.Date, &i.CreatedAt)
		if err != nil {
			return []Income{}, err
		}
		income = append(income, i)
	}
	if err = rows.Err(); err != nil {
		return []Income{}, err
	}
	return income, nil
}

func GetId(db *sql.DB, id int64) (Income, error) {
	income, err := Get(db)
	if err != nil {
		return Income{}, err
	}
	for _, i := range income {
		if i.ID == id {
			return i, nil
		}
	}
	return Income{}, errors.New("user does not Exist")
}

func (i *Income) PutValues(db *sql.DB) (int64, error) {
	query := `UPDATE income SET amount=$1, category=$2, source=$3, date=$4 WHERE id=$5`
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

func Delete(db *sql.DB, id int64) error {
	query := `DELETE FROM income where id=$1`
	_, err := db.Exec(query, id)
	return err
}
