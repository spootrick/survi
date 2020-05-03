package routes

import (
	"github.com/spootrick/survi/api/controller"
	"net/http"
)

var userDetailRoutes = []Route{
	{
		URI:          "/users/{id}/details",
		Method:       http.MethodGet,
		Handler:      controller.GetUserDetail,
		AuthRequired: true,
	},
	{
		URI:          "/users/details",
		Method:       http.MethodPost,
		Handler:      controller.CreateUserDetail,
		AuthRequired: true,
	},
	{
		URI:          "/users/{id}/details",
		Method:       http.MethodPut,
		Handler:      controller.UpdateUserDetail,
		AuthRequired: true,
	},
	{
		URI:          "/users/{id}/details",
		Method:       http.MethodDelete,
		Handler:      nil,
		AuthRequired: true,
	},
}
