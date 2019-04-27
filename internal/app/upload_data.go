package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Telemetry(w http.ResponseWriter, r *http.Request) {
	deviceInfo, _ := r.Context().Value("deviceInfo").(map[string]string)
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		body := map[string]string{"error": err.Error()}
		_ = json.NewEncoder(w).Encode(body)
		return
	}
	TelemetryData(deviceInfo, data)

}

func TelemetryFile(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
