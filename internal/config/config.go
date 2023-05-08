package config

import (
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	InProduction bool
	Session      *scs.SessionManager
}

// JwtSecret is a constant that holds the secret key for signing JWT tokens.
const JwtSecret = "test124"
