package middleware

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   App.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	// Exclude the routes from CSRF protection.
	csrfHandler.ExemptPath("/login")
	csrfHandler.ExemptPath("/set-session-value")
	csrfHandler.ExemptPath("/get-session-value")

	return csrfHandler
}
