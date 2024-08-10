package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kentlouisetonino/gointegration/src/libs"
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

	// Log the request.
	libs.APILogger("/api/health", response.StatusCode)
}
