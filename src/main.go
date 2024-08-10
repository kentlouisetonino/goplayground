package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kentlouisetonino/gointegration/src/handlers"
)

func main() {
	// Get the PORT from the environment variable.
	port := os.Getenv("PORT")
	if port == "" {
		port = "11000"
	}

	// Health check.
	http.HandleFunc("/api/health", handlers.Health)

	// Start the server
	log.Printf("Server is listening on port %s.", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
