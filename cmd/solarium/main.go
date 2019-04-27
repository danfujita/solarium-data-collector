package main

import (
    "fmt"
    "net/http"
    "solarium-golang/internal/app"
)


func main() {
    http.HandleFunc("/api/telemetry", app.Telemetry)
    http.HandleFunc("/api/telemetry_file", app.Telemetry_file)
    fmt.Print("Service starting at port 8080")
    http.ListenAndServe(":8080", nil)
}

