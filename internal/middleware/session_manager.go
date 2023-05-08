package middleware

import (
	"net/http"
)

func SessionLoad(next http.Handler) http.Handler {
	return Session.LoadAndSave(next)
}
