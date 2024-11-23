package utils

import (
	"log"
	"time"
)

// DoCleanup simulates cleanup tasks during server shutdown
func DoCleanup() {
	log.Println("Performing cleanup tasks...")
	time.Sleep(2 * time.Second) // Simulate cleanup delay
	log.Println("Cleanup completed.")
}

