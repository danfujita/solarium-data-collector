package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"solarium-data-collector/internal/app"
	"solarium-data-collector/internal/config_reader"
	"solarium-data-collector/internal/middleware"
	"strconv"
)

func main() {
	r := chi.NewRouter()
	myConfig := config_reader.Config()
	if myConfig.PortNumber == 0 {
		myConfig.PortNumber = 8080
	}
	strPortNumber := strconv.Itoa(myConfig.PortNumber)
	r.Use(middleware.Middleware)
	r.Post("/api/telemetry", app.Telemetry)
	r.Post("/api/payload", app.Payload)
	fmt.Print("Service starting at port " + strPortNumber)
	http.ListenAndServe(":"+strPortNumber, r)
}
