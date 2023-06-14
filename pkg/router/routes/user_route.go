package routes

import (
	"Api-Go/pkg/handlers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               handlers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               handlers.GetUsers,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               handlers.GetUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodPut,
		Function:               handlers.UpdateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodDelete,
		Function:               handlers.RemoveUser,
		RequiresAuthentication: false,
	},
}

func init() {
	Routes = append(Routes, userRoutes...)
}
