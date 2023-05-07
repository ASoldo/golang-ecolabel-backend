// Package handlers provides HTTP handlers for the golang-ecolabel-backend project.
package handlers

import (
	"encoding/json"
	"net/http"
)

// jsonResponse is a generic JSON response structure used to return messages in the response body.
type jsonResponse struct {
	Message string `json:"message"`
}

// HandleDashboard is an HTTP handler that returns a welcome message for the dashboard.
// This handler requires JWT authentication and should be protected by the JwtAuthMiddleware.
func HandleDashboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := jsonResponse{
		Message: "Welcome to the dashboard",
	}
	json.NewEncoder(w).Encode(response)
}
