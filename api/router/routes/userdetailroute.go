package routes

import (
	"github.com/spootrick/survi/api/controller"
	"net/http"
)

var userDetailRoutes = []Route{
	{
		URI:     "/users/{id}/details",
		Method:  http.MethodGet,
		Handler: controller.GetUserDetail,
	},
	{
		URI:     "/users/details",
		Method:  http.MethodPost,
		Handler: controller.CreateUserDetail,
	},
	{
		URI:     "/users/{id}/details",
		Method:  http.MethodPut,
		Handler: nil,
	},
	{
		URI:     "/users/{id}/details",
		Method:  http.MethodDelete,
		Handler: nil,
	},
}
