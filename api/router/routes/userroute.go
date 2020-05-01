package routes

import (
	"github.com/spootrick/survi/api/controller"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:     "/users",
		Method:  http.MethodGet,
		Handler: controller.GetAllUsers,
	},
	{
		URI:     "/users",
		Method:  http.MethodPost,
		Handler: controller.CreateUser,
	},
	{
		URI:     "/users/{id}",
		Method:  http.MethodGet,
		Handler: controller.GetUser,
	},
	{
		URI:     "/users/{id}",
		Method:  http.MethodPut,
		Handler: controller.UpdateUser,
	},
	{
		URI:     "/users/{id}",
		Method:  http.MethodDelete,
		Handler: controller.DeleteUser,
	},
}
