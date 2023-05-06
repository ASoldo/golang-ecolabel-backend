package handlers

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Message string `json:"message"`
}

func HandleDashboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := jsonResponse{
		Message: "Welcome to the dashboard",
	}
	json.NewEncoder(w).Encode(response)
}
