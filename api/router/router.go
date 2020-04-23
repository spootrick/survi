package router

import (
	"github.com/gorilla/mux"
	"github.com/spootrick/survi/api/router/routes"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutesWithMiddleware(r)
}
