// Package middleware provides custom middleware functions for the golang-ecolabel-backend project.
package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"github.com/ASoldo/golang-ecolabel-backend/internal/config"
)

// JwtAuthMiddleware is an HTTP middleware function that checks for the presence and validity of a JWT token in the
// Authorization header of incoming HTTP requests. If the token is valid, the request is allowed to continue.
// Otherwise, an HTTP error is returned.
func JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for the presence of an Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Check the format of the Authorization header
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}

		// Parse the JWT token
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Invalid signing method")
			}
			return []byte(config.JwtSecret), nil
		})

		// Check for errors and the validity of the token
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// If the token is valid, continue processing the request
		next.ServeHTTP(w, r)
	})
}
