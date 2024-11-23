package handlers

import (
	"net/http"
)

// HandleLiveness handles the liveness probe
func HandleLiveness() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is still healthy (alive)."))
	})
}

