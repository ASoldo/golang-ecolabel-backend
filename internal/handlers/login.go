package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ASoldo/golang-ecolabel-backend/internal/services"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// userService is an instance of the UserServiceImpl struct
var userService services.UserService = services.NewUserService()

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var loginReq loginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := userService.Authenticate(loginReq.Username, loginReq.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
