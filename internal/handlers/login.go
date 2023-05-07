// Package handlers provides HTTP handlers for the golang-ecolabel-backend project.
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ASoldo/golang-ecolabel-backend/internal/services"
)

// loginRequest represents the expected structure of the login request's JSON body.
type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// userService is an instance of the UserServiceImpl struct.
var userService services.UserService = services.NewUserService()

// HandleLogin is an HTTP handler that processes login requests.
// It reads a JSON request body containing a username and password,
// authenticates the credentials using the userService, and returns a JSON response with a JWT token if the authentication is successful.
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
