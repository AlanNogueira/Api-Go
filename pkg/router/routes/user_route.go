package routes

import (
	"Api-Go/pkg/handlers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                    "/user",
		Method:                 http.MethodPost,
		Function:               handlers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/user",
		Method:                 http.MethodGet,
		Function:               handlers.GetUsers,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/user/authenticate",
		Method:                 http.MethodPost,
		Function:               handlers.AuthenticateUser,
		RequiresAuthentication: false,
	},
}

func init() {
	Routes = append(Routes, userRoutes...)
}
