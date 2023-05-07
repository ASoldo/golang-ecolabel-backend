// Package models contains the data structures and functions for managing users and generating JWT tokens.
package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// User represents a user with a username and password.
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// DemoUser is a pre-defined user with a username and password for demonstration purposes.
var DemoUser = User{
	Username: "test",
	Password: "test",
}

// GenerateToken creates a JWT token for the DemoUser with a 1-hour expiration time.
// It returns the signed token string.
func GenerateToken(secret string) string {
	// Create a new JWT token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": DemoUser.Username,
		"exp":      time.Now().Add(1 * time.Hour).Unix(),
	})

	// Sign the token using the provided secret
	tokenString, _ := token.SignedString([]byte(secret))

	// Return the signed token string
	return tokenString
}
