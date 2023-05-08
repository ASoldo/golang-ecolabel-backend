// Package routes provides the API routing setup for the golang-ecolabel-backend project.
package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ASoldo/golang-ecolabel-backend/internal/handlers"
	customMiddleware "github.com/ASoldo/golang-ecolabel-backend/internal/middleware"
)

// SetupRoutes configures the API routes and returns the configured router.
// It sets up common middlewares and defines routes for the login and dashboard endpoints.
func SetupRoutes() *chi.Mux {
	// Create a new Chi router
	r := chi.NewRouter()

	// Set up common middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// r.Use(customMiddleware.SessionLoad)
	r.Use(customMiddleware.NoSurf)

	// Define the /login route with a POST method
	r.Post("/login", handlers.HandleLogin)

	// Define the /dashboard route with a GET method and JWT authentication middleware
	r.With(customMiddleware.JwtAuthMiddleware).Get("/dashboard", handlers.HandleDashboard)

	// Define the /set-session-value route with a POST method
	// r.Post("/set-session-value", handlers.HandleSetSessionValue)

	// Define the /get-session-value route with a GET method
	// r.Get("/get-session-value", handlers.HandleGetSessionValue)

	// Define the /csrf-token route with a GET method
	r.Get("/csrf-token", handlers.HandleGetCSRFToken)

	// Define the /submit route with NoSurf check for csrf-token
	r.With(customMiddleware.NoSurf).Post("/submit", handlers.HandleSubmit)

	// Return the configured router
	return r
}
