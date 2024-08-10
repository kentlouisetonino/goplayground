package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	var response HealthResponse

	// Ensure the request is a GET method.
	if r.Method != http.MethodGet {
		response = HealthResponse{
			Message:    "Method not allowed",
			StatusCode: http.StatusMethodNotAllowed,
		}
	} else {
		response = HealthResponse{
			Message:    "Server is ready.",
			StatusCode: http.StatusOK,
		}
	}

	// Set the content type header.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	// Encode the response as JSON and send it.
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Colors for formatting
	blue := "\033[34m"
	green := "\033[32m"
	red := "\033[31m"
	reset := "\033[0m"

	// Color the status code based on its value.
	statusColor := green
	if response.StatusCode >= 400 && response.StatusCode <= 500 {
		statusColor = red
	}

	// Log the request.
	log.Printf("%s/api/health%s %s%d%s", blue, reset, statusColor, response.StatusCode, reset)
}
