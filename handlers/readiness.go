package handlers

import (
	"context"
	"net/http"
)

// HandleReadiness handles the readiness probe
func HandleReadiness(ctx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-ctx.Done(): // Server is shutting down
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(" Server is shutting down. (Unable to handle requests anymore)"))
		default: // Server is ready
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Server is ready"))
		}
	})
}
