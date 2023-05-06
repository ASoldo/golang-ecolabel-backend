package main

// "github.com/ASoldo/golang-ecolabel-backend/internal/errors"
import (
	"net/http"

	"github.com/ASoldo/golang-ecolabel-backend/internal/logger"
	"github.com/ASoldo/golang-ecolabel-backend/internal/routes"
)

func main() {
	// err := errors.NewAppError(400, "Bad Request")
	// logger.Error.Println(err)

	r := routes.SetupRoutes()
	logger.Info.Println("Application started")
	http.ListenAndServe(":3000", r)
}
