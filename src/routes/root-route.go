package routes

import (
	"net/http"

	"github.com/kentlouisetonino/go-integration/src/controllers"
)

func RootRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet && r.URL.Path == "/" {
		controllers.RootController(w)
	}
}
