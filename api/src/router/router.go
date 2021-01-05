package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate returns a configures mux router
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigureRoutes(r)
}
