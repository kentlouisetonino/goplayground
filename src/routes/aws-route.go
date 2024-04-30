package routes

import (
	"fmt"
	"net/http"
)

func AWSRoute(w http.ResponseWriter, r *http.Request) {
	// api/aws/s3/upload
	if r.Method == http.MethodPost && r.URL.Path == "/s3/upload" {
		fmt.Fprintf(w, "Hello Hell!")
	}
}
