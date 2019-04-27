package app

import (
	"fmt"
	"net/http"
)

func Telemetry(w http.ResponseWriter, r *http.Request) {
	deviceInfo, _ := r.Context().Value("token").(map[string]string)
    print(deviceInfo)
}
func TelemetryFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
