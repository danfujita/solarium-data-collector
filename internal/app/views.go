package app

import (
	"bytes"
	"encoding/json"
	"io"
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
	mReader, err := r.MultipartReader()
	deviceInfo, _ := r.Context().Value("deviceInfo").(map[string]string)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		body := map[string]string{"error": err.Error()}
		_ = json.NewEncoder(w).Encode(body)
		return
	}
	for {
		part, err := mReader.NextPart()
		if err != nil {
			break
		}
		if part.FormName() == "data" {
			buf := bytes.NewBuffer(nil)
			io.Copy(buf, part)
			err := AddFileToS3(buf.Bytes(), deviceInfo["device_id"])
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
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	body := map[string]string{"error": "Data key was not found."}
	_ = json.NewEncoder(w).Encode(body)
	return
}
