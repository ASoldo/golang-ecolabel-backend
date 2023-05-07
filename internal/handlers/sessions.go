package handlers

import (
	"net/http"

	"github.com/ASoldo/golang-ecolabel-backend/internal/middleware"
)

// HandleSetSessionValue sets a value in the session.
func HandleSetSessionValue(w http.ResponseWriter, r *http.Request) {
	middleware.Session.Put(r.Context(), "session-key", "This is my Session")
	w.Write([]byte("Session value set"))
}

// HandleGetSessionValue retrieves a value from the session.
func HandleGetSessionValue(w http.ResponseWriter, r *http.Request) {
	value, ok := middleware.Session.Get(r.Context(), "session-key").(string)
	if !ok {
		http.Error(w, "Error getting session value", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Session value: " + value))
}
