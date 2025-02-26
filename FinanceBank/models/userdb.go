package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            string    `json:"id"`
	Account_id    string    `json:"account_id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Password_hash string    `json:"password_hash"`
	Currency      string    `json:"currency"`
	CreatedAt     time.Time `json:"created_at"`
}

func (u *User) Create(db *sql.DB) (*User, error) {
	query := `
	INSERT INTO users (id,account_id,name, email, password_hash, currency) 
	VALUES ($1, $2, $3, $4,$5,$6) 
	RETURNING id, created_at;
	`
	u.ID = uuid.NewString()
	u.Account_id = uuid.NewString()
	err := db.QueryRow(query, u.ID, u.Account_id, u.Name, u.Email, u.Password_hash, u.Currency).Scan(&u.ID, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	query = `INSERT INTO accounts(id,user_id,amount) 
	VALUES ($1,$2,$3)`
	_, err = db.Exec(query, u.Account_id, u.ID, 0)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUsers(db *sql.DB) ([]User, error) {
	var income []User
	query := `
	SELECT id,account_id,name,email,currency,created_at FROM users
	`
	rows, err := db.Query(query)
	if err != nil {
		return []User{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Account_id, &u.Name, &u.Email, &u.Currency, &u.CreatedAt)
		if err != nil {
			return []User{}, err
		}
		income = append(income, u)
	}
	if err = rows.Err(); err != nil {
		return []User{}, err
	}
	return income, nil
}

func GetIdUser(db *sql.DB, id string) (User, error) {
	var u User
	query := `SELECT id, account_id, name, email, currency, created_at FROM users WHERE id=$1`
	err := db.QueryRow(query, id).Scan(&u.ID, &u.Account_id, &u.Name, &u.Email, &u.Currency, &u.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, errors.New("does not exist")
		}
		return User{}, err
	}
	return u, nil
}

func (u *User) PutValues(db *sql.DB) (int64, error) {
	query := `UPDATE users SET name=$1, email=$2, currency=$3 WHERE id=$4`
	result, err := db.Exec(query, u.Name, u.Email, u.Currency, u.ID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func DeleteUser(db *sql.DB, id string) error {
	query := `DELETE FROM users where id=$1 CASCADE`
	_, err := db.Exec(query, id)
	return err
}

func CheckExistAndDifferent(db *sql.DB, id1 string, id2 string) error {
	if id1 == id2 {
		return errors.New("transaction not possible with same UserID's")
	}

	query := `SELECT COUNT(*) FROM users WHERE id IN ($1, $2)`
	var count int
	err := db.QueryRow(query, id1, id2).Scan(&count)
	if err != nil {
		return err
	}

	if count != 2 {
		return errors.New("one or both user IDs not found")
	}

	return nil
}

