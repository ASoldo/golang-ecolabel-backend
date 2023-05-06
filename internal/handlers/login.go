package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/ASoldo/golang-ecolabel-backend/internal/config"
	"github.com/ASoldo/golang-ecolabel-backend/internal/models"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var loginReq loginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if loginReq.Username == models.DemoUser.Username &&
		loginReq.Password == models.DemoUser.Password {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": loginReq.Username,
			"exp":      time.Now().Add(1 * time.Hour).Unix(),
		})

		tokenString, err := token.SignedString([]byte(config.JwtSecret))
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"token": tokenString,
		})
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}
