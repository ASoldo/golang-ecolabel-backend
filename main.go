package main

import (
	"net/http"

	"github.com/ASoldo/golang-ecolabel-backend/internal/routes"
)

func main() {
	r := routes.SetupRoutes()
	http.ListenAndServe(":3000", r)
}
