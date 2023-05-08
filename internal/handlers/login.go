// Package handlers provides HTTP handlers for the golang-ecolabel-backend project.
package handlers

import (
	"encoding/json"
	"fmt"
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
		fmt.Printf("Invalid request body: %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if loginReq.Username == "" || loginReq.Password == "" {
		fmt.Println("Empty username or password")
		http.Error(w, "Empty username or password", http.StatusBadRequest)
		return
	}

	token, err := userService.Authenticate(loginReq.Username, loginReq.Password)
	if err != nil {
		fmt.Printf("Invalid credentials: %v\n", err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
