package routes

import (
	"Api-Go/pkg/handlers"
	"net/http"
)

var rentalCompanyRoutes = []Route{
	{
		URI:                    "/rentalCompanies",
		Method:                 http.MethodPost,
		Function:               handlers.CreateRentalCompany,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/rentalCompanies",
		Method:                 http.MethodGet,
		Function:               handlers.CreateRentalCompany,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/rentalCompanies/authenticate",
		Method:                 http.MethodPost,
		Function:               handlers.AuthenticateUser,
		RequiresAuthentication: false,
	},
}

func init() {
	Routes = append(Routes, rentalCompanyRoutes...)
}
