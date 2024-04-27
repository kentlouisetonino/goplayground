package routes

import (
	"net/http"

	"github.com/kentlouisetonino/go-integration/src/controllers"
)

func RootRoute(w http.ResponseWriter, r *http.Request) {
	controllers.RootController()
}
