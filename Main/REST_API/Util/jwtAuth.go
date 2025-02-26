package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secret"

func GetToken(email string, id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(time.Minute * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected Signing Method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("could Not Parse Token")
	}
	validToken := parsedToken.Valid
	if !validToken {
		return 0, errors.New("invalid token")
	}

	claim, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("unexpected Signing Method")
	}
	id := claim["id"].(float64)
	return int64(id), nil
}
