package main

import (
	"net/http"
)

// Server Readiness Handler
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, map[string]string{"message": "Ready"})
}
