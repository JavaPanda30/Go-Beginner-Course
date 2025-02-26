package models

import (
	"database/sql"
	"errors"
	"fmt"
	"example.com/financetracker/db"
)

type Account struct {
	Id     string
	UserID string
	Amount float64
}

func UpdateAccountAmt(amount float64, userID string) error {
	query := `UPDATE accounts SET amount = amount + $1 WHERE user_id = $2`
	result, err := db.DB.Exec(query, amount, userID)
	if err != nil {
		return fmt.Errorf("failed to update account balance: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("no account found for the given user_id")
	}
	return nil
}

func GetAccountAmount(userID string) (float64, error) {
	query := `SELECT amount FROM accounts WHERE user_id = $1`
	var amt float64
	err := db.DB.QueryRow(query, userID).Scan(&amt)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("no account found for the given user_id")
		}
		return 0, fmt.Errorf("error fetching account balance: %w", err)
	}

	return amt, nil
}
