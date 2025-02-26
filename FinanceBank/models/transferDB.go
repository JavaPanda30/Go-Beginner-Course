package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        string    `json:"id"`
	From_user string    `json:"from_user"`
	To_user   string    `json:"to_user"`
	Amount    float64   `json:"amount"`
	Time      time.Time `json:"time"`
}

func (t *Transaction) Insert(db *sql.DB) (*Transaction, error) {
	err := CheckExistAndDifferent(db, t.From_user, t.To_user)
	if err != nil {
		return nil, err
	}
	_, err = GetIdUser(db, t.From_user)
	if err != nil {
		return nil, err
	}
	amt, err := GetAccountAmount(t.From_user)
	if err != nil {
		return &Transaction{}, err
	}
	UpdateAccountAmt(amt-t.Amount, t.From_user)
	amt, err = GetAccountAmount(t.To_user)
	if err != nil {
		return &Transaction{}, err
	}
	UpdateAccountAmt(amt+t.Amount, t.To_user)
	query := `INSERT INTO transfers(id,from_user,to_user,amount,time) 
	VALUES ($1,$2,$3,$4,$5) 
	RETURNING from_user,to_user,id`
	t.ID = uuid.NewString()
	err = db.QueryRow(query, t.ID, t.From_user, t.To_user, t.Amount, t.Time).Scan(&t.From_user, &t.To_user, &t.ID)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func GetTransactionById(db *sql.DB, id string) (*Transaction, error) {
	var transaction Transaction
	query := `SELECT id, from_user, to_user, amount, time FROM transfers WHERE id = $1`
	err := db.QueryRow(query, id).Scan(
		&transaction.ID,
		&transaction.From_user,
		&transaction.To_user,
		&transaction.Amount,
		&transaction.Time,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("transaction not found")
		}
		return nil, err
	}
	return &transaction, nil
}
