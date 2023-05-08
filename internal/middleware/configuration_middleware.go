package middleware

import (
	"log"

	"github.com/alexedwards/scs/v2"

	"github.com/ASoldo/golang-ecolabel-backend/internal/config"
)

var (
	App      config.AppConfig
	Session  *scs.SessionManager
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)
