package main

import "net/http"

// Sends JSON error response to http client
func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong!")
}
