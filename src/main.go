package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kentlouisetonino/gointegration/src/handlers"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading env variables.")
	}

	// Get the PORT from the environment variable.
	port := os.Getenv("PORT")

	// Health check.
	http.HandleFunc("/api/health", handlers.Health)

	// Start the server
	log.Printf("Server is listening on port %s.", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
