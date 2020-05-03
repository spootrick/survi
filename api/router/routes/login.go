package routes

import (
	"github.com/spootrick/survi/api/controller"
	"net/http"
)

var loginRoutes = []Route{
	{
		URI:          "/login",
		Method:       http.MethodPost,
		Handler:      controller.Login,
		AuthRequired: false,
	},
}
