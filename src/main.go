package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kentlouisetonino/go-labs/src/handlers"
	"github.com/kentlouisetonino/go-labs/src/libs"
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

	// Log the available api endpoints.
	libs.AddNewLine()
	libs.APIEndpointLogger("/api/health")

	// Start the server.
	log.Printf("%sServer is listening on port %s.%s", libs.Green, port, libs.Reset)

	// Handle the error.
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
