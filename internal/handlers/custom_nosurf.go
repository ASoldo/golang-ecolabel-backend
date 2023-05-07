package handlers

import (
	"net/http"
)

func HandleSubmit(w http.ResponseWriter, r *http.Request) {
	// Process the form submission and perform any necessary actions
	w.Write([]byte("Form submitted successfully"))
}
