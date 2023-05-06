package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ASoldo/golang-ecolabel-backend/internal/handlers"
	customMiddleware "github.com/ASoldo/golang-ecolabel-backend/internal/middleware"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/login", handlers.HandleLogin)
	r.With(customMiddleware.JwtAuthMiddleware).Get("/dashboard", handlers.HandleDashboard)

	return r
}
