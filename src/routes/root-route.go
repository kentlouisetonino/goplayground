package routes

import (
	"net/http"

	"github.com/kentlouisetonino/go-integration/src/controllers"
)

type Response struct {
	Message string
}

func RootRoute(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		controllers.RootController(w)
	}
}
