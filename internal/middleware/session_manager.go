package middleware

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
)

var Session *scs.SessionManager

func SessionLoad(next http.Handler) http.Handler {
	return Session.LoadAndSave(next)
}
