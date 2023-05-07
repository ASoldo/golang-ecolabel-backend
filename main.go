// Package main is the entry point for the golang-ecolabel-backend application.
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"

	"github.com/ASoldo/golang-ecolabel-backend/internal/logger"
	"github.com/ASoldo/golang-ecolabel-backend/internal/middleware"
	"github.com/ASoldo/golang-ecolabel-backend/internal/routes"
)

// main is the entry point of the golang-ecolabel-backend application.
// It sets up the routes, logs that the application has started, and starts the HTTP server.
func main() {
	middleware.Session = scs.New()
	middleware.Session.Lifetime = 24 * time.Hour
	middleware.Session.Cookie.Persist = true
	middleware.Session.Cookie.SameSite = http.SameSiteLaxMode
	middleware.Session.Cookie.Secure = false
	fmt.Println(middleware.Session)
	// Set up the routes
	r := routes.SetupRoutes()

	// Log that the application has started
	logger.Info.Println("Application started")

	// Start the HTTP server on port 3000
	http.ListenAndServe(":3000", r)
}
