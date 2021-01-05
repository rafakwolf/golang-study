package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route representa uma rota
type Route struct {
	URI          string
	Method       string
	Handler      func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

// ConfigureRoutes configure and returns all routes
func ConfigureRoutes(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}

	return r
}
