package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kentlouisetonino/go-labs/src/libs"
	"github.com/kentlouisetonino/go-labs/src/models"
)

func Health(w http.ResponseWriter, r *http.Request) {
	var response models.HTTPResponse

	// Ensure the request is a GET method.
	if r.Method != http.MethodGet {
		response = models.HTTPResponse{
			Message:    "Method not allowed",
			StatusCode: http.StatusMethodNotAllowed,
		}
	} else {
		response = models.HTTPResponse{
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
