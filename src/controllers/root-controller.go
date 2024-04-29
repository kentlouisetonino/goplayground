package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string
}

func RootController(w http.ResponseWriter) {
	jsonResponse, err := json.Marshal(Response{Message: "hello"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(jsonResponse))
}
