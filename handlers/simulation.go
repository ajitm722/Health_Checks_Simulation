package handlers

import (
	"fmt"
	"net/http"
)

// HandleSimulationPage serves a simple HTML interface for simulation
func HandleSimulationPage() http.Handler {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Server Simulation</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
        }
        .endpoint {
            margin: 10px 0;
        }
    </style>
</head>
<body>
    <h1>Server Simulation</h1>
    <p>This page demonstrates readiness and liveness probes for a cloud-native service.</p>
    <div class="endpoint">
        <h2>Endpoints</h2>
        <ul>
            <li><strong>Readiness Probe:</strong> <a href="/ready" target="_blank">/ready</a></li>
            <li><strong>Liveness Probe:</strong> <a href="/health" target="_blank">/health</a></li>
        </ul>
    </div>
    <div class="instructions">
        <h2>Instructions</h2>
        <ol>
            <li>Access the <a href="/ready" target="_blank">/ready</a> endpoint to check if the server is ready to handle requests.</li>
            <li>Access the <a href="/health" target="_blank">/health</a> endpoint to check if the server is alive.</li>
            <li>Simulate a shutdown by sending a termination signal (e.g., press <code>Ctrl+C</code> in the terminal).</li>
            <li>Observe how the readiness probe returns a <code>503</code> status during shutdown.</li>
            <li>Observe how cleanup tasks are logged, and the server shuts down gracefully.</li>
        </ol>
    </div>
</body>
</html>`
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, html)
	})
}

