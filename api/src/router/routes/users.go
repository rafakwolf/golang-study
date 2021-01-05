package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Handler:      controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Handler:      controllers.ListUsers,
		AuthRequired: false,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodGet,
		Handler:      controllers.GetUser,
		AuthRequired: false,
	},
	{
		URI:          "/users/{userID}",
		Method:       http.MethodPut,
		Handler:      controllers.UpdateUser,
		AuthRequired: false,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodDelete,
		Handler:      controllers.RemoveUser,
		AuthRequired: false,
	},
}
