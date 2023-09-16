package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Sends error response as JSON object to http client
func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Internal server error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponse{Error: msg})
}

// Sends response as JSON object to http client.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
