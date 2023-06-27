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
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               handlers.GetUsers,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{taxNumber}",
		Method:                 http.MethodGet,
		Function:               handlers.GetUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodPut,
		Function:               handlers.UpdateUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodDelete,
		Function:               handlers.RemoveUser,
		RequiresAuthentication: true,
	},
}

func init() {
	Routes = append(Routes, userRoutes...)
}
