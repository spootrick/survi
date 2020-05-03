package routes

import (
	"github.com/spootrick/survi/api/controller"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Handler:      controller.GetAllUsers,
		AuthRequired: true,
	},
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Handler:      controller.CreateUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{id}",
		Method:       http.MethodGet,
		Handler:      controller.GetUser,
		AuthRequired: false,
	},
	{
		URI:          "/users/{id}",
		Method:       http.MethodPut,
		Handler:      controller.UpdateUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{id}",
		Method:       http.MethodDelete,
		Handler:      controller.DeleteUser,
		AuthRequired: true,
	},
}
