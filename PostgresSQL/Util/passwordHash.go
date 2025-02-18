package util

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	var salt = 10
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CheckPassword(hash,password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
