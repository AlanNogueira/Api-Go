package routes

import (
	"Api-Go/pkg/handlers"
	"net/http"
)

var rentRoutes = []Route{
	{
		URI:                    "/rents",
		Method:                 http.MethodPost,
		Function:               handlers.CreateRent,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/rents",
		Method:                 http.MethodGet,
		Function:               handlers.GetNotDeliveredRents,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/rents/{clientName}",
		Method:                 http.MethodGet,
		Function:               handlers.GetRent,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/rents/finalize/{rentId}",
		Method:                 http.MethodPut,
		Function:               handlers.FinalizeRent,
		RequiresAuthentication: true,
	},
}

func init() {
	Routes = append(Routes, rentRoutes...)
}
