package handlers

import (
	"encoding/json"
	"net/http"

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
		tokenString := models.GenerateToken(config.JwtSecret)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"token": tokenString,
		})
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}
