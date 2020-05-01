package routes

import (
	"github.com/gorilla/mux"
	"github.com/spootrick/survi/api/middleware"
	"net/http"
)

type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func Load() []Route {
	routes := userRoutes

	return routes
}

// without middlewares
func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
	return r
}

func SetupRoutesWithMiddleware(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.URI,
			middleware.SetMiddlewareLogger(
				middleware.SetMiddlewareJSON(route.Handler)),
		).Methods(route.Method)
	}
	return r
}
