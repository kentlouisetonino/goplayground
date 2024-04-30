package routes

import (
	"fmt"
	"net/http"
)

func AWSRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if r.Method == http.MethodPost && r.URL.Path == "/s3/upload" {
		fmt.Fprintf(w, "Hello Hell!")
	}
}
