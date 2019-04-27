package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"solarium-golang/internal/app"
	"solarium-golang/internal/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Middleware)
	r.Post("/api/telemetry", app.Telemetry)
	r.Post("/api/payload", app.Payload)
	fmt.Print("Service starting at port 8080")
	http.ListenAndServe(":8080", r)
}
