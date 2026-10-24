package main

import (
	"net/http"
)

// Error Handler
func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went Wrong")
}
