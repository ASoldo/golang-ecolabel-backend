package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var DemoUser = User{
	Username: "test",
	Password: "test",
}

func GenerateToken(secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": DemoUser.Username,
		"exp":      time.Now().Add(1 * time.Hour).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(secret))

	return tokenString
}
