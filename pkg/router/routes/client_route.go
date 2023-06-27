package routes

import (
	"Api-Go/pkg/handlers"
	"net/http"
)

var clientRoutes = []Route{
	{
		URI:                    "/clients",
		Method:                 http.MethodPost,
		Function:               handlers.CreateClient,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/clients",
		Method:                 http.MethodGet,
		Function:               handlers.GetClients,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/clients/{taxNumber}",
		Method:                 http.MethodGet,
		Function:               handlers.GetClient,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/clients/{clientId}",
		Method:                 http.MethodPut,
		Function:               handlers.UpdateClient,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/clients/{clientId}",
		Method:                 http.MethodDelete,
		Function:               handlers.RemoveClient,
		RequiresAuthentication: true,
	},
}

func init() {
	Routes = append(Routes, clientRoutes...)
}
