package models

import (
	"errors"

	util "example.com/eventbook/Util"
	"example.com/eventbook/db"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (user *User) Save() error {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	query := `
    INSERT INTO users (name, email, password) 
    VALUES ($1, $2, $3)
    RETURNING id`
	err = db.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers() ([]User, error) {
	query := `SELECT * FROM users`
	row, err := db.DB.Query(query)
	if err != nil {
		return []User{}, err
	}
	defer row.Close()
	var users []User
	for row.Next() {
		var user User
		err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
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
		err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
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
	query := `DELETE FROM users WHERE id=$1 RETURNING id, name, email, password`
	var user User
	err := db.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id,password FROM users WHERE email=$1`
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