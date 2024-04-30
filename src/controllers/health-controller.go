package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func HealthController(w http.ResponseWriter) {
	jsonResponse, err := json.Marshal(Response{Message: "Server is running.", Code: http.StatusOK})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(jsonResponse))
}
