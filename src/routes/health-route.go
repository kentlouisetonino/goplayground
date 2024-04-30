package routes

import (
	"net/http"

	"github.com/kentlouisetonino/go-integration/src/controllers"
)

func HealthRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		controllers.HealthController(w)
		return
	}

	http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
}
