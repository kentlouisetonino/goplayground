package routes

import (
	"fmt"
	"net/http"
)

func AWSS3Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println(r.URL.Path)
		return
	}

	http.Error(w, "Endpoint not found.", http.StatusNotFound)
}
