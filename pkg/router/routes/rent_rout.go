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
		RequiresAuthentication: false,
	},
	{
		URI:                    "/rents",
		Method:                 http.MethodGet,
		Function:               handlers.GetNotDeliveredRents,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/rents/{userName}",
		Method:                 http.MethodGet,
		Function:               handlers.GetRent,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/rents/finalize/{rentId}",
		Method:                 http.MethodPut,
		Function:               handlers.FinalizeRent,
		RequiresAuthentication: false,
	},
}

func init() {
	Routes = append(Routes, rentRoutes...)
}
