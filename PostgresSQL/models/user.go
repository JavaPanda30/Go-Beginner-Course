package models

import (
	"errors"

	util "example.com/eventbook/Util"
	"example.com/eventbook/db"
)

type User struct {
	ID       int64  `json:"id"`
	UserName string `json:"username"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (u *User) Save() error {
	query := `INSERT INTO user(name,email,password) VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	u.Password, _ = util.HashPassword(u.Password)
	defer stmt.Close()
	res, err := stmt.Exec(u.UserName, u.Email, u.Password)
	if err != nil {
		return err
	}
	u.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers() ([]User, error) {
	query := `SELECT * FROM user`
	row, err := db.DB.Query(query)
	if err != nil {
		return []User{}, err
	}
	defer row.Close()
	var users []User
	for row.Next() {
		var user User
		err := row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserById(id int64) (User, error) {
	query := `SELECT * FROM user`
	row, err := db.DB.Query(query)
	if err != nil {
		return User{}, err
	}
	defer row.Close()
	for row.Next() {
		var user User
		err := row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password)
		if err != nil {
			return User{}, err
		}
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("UserID Not Found")
}

func Delete(id int64) (User, error) {
	query := `DELETE FROM user WHERE id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return User{}, err
	}
	user, err := GetUserById(id)
	if err != nil {
		return User{}, err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(id); err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id,password FROM user WHERE email=?`
	row := db.DB.QueryRow(query, u.Email)
	var hashpassword string
	err := row.Scan(&u.ID, &hashpassword)
	if err != nil {
		return err
	}
	err = util.CheckPassword(hashpassword, u.Password)
	if err != nil {
		return errors.New("email or Password Does not match")
	}
	return nil
}

// func (u *User) GetEventsCreatedByUserId() ([]Event, error) {
// 	// query=`SELECT * FROM events WHERE userId=?`

// }
