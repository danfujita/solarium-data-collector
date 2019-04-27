package main

import (
	"fmt"
	"net/http"
	"solarium-golang/internal/app"
	"solarium-golang/internal/middleware"
)

func main() {
	http.Handle("/api/telemetry", middleware.Middleware(http.HandlerFunc(app.Telemetry)))
	http.Handle("/api/telemetry_file", middleware.Middleware(http.HandlerFunc(app.TelemetryFile)))
	fmt.Print("Service starting at port 8080")
	http.ListenAndServe(":8080", nil)
}
