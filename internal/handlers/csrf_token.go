package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/justinas/nosurf"
)

type csrfTokenResponse struct {
	CSRFToken string `json:"csrfToken"`
}

func HandleGetCSRFToken(w http.ResponseWriter, r *http.Request) {
	token := nosurf.Token(r)

	response := csrfTokenResponse{
		CSRFToken: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
