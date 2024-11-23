package main

import (
	"context"
	"health_checks/handlers"
	"health_checks/utils"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create a context that listens for termination signals (SIGTERM, SIGINT, SIGQUIT)
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT,
	)
	defer cancel() // Ensure that the context is canceled on program exit

	// Set up the HTTP server
	server := &http.Server{Addr: ":8080"}

	// Register HTTP handlers
	http.Handle("/", handlers.HandleSimulationPage())    // Serve simulation HTML page
	http.Handle("/ready", handlers.HandleReadiness(ctx)) // Readiness probe
	http.Handle("/health", handlers.HandleLiveness())    // Liveness probe

	// Start a goroutine to handle graceful shutdown
	go func() {
		<-ctx.Done()
		log.Println("Got shutdown signal. Shutting down gracefully...")

		// Perform cleanup tasks
		utils.DoCleanup()

		// Simulate a readiness delay before actual shutdown
		<-time.After(8 * time.Second)

		// Shutdown the HTTP server gracefully
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("Error while stopping HTTP listener: %s", err)
		}
	}()

	// Start the HTTP server
	log.Println("Starting server on :8080")
	log.Println(server.ListenAndServe())
}
