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

	err = TelemetryData(deviceInfo, data)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		body := map[string]string{"error": err.Error()}
		_ = json.NewEncoder(w).Encode(body)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		body := map[string]bool{"success": true}
		_ = json.NewEncoder(w).Encode(body)
		return
	}
}

func Payload(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
