package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, status int, msg string) {
	if status > 499 {
		log.Printf("Response with error 5xx: %s", msg)
	}

	type ErrorResponse struct {
		ErrorResponse string `json:"error"`
	}
	responseWithJSON(w, status, ErrorResponse{
		ErrorResponse: msg,
	})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	resp, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal JSON response: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}
