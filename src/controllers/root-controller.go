package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func RootController(w http.ResponseWriter) {
	jsonResponse, err := json.Marshal(Response{Message: "Meaningless Life!"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(jsonResponse))
}
