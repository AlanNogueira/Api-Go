package routes

import (
	"Api-Go/pkg/handlers"
	"net/http"
)

var reportsRouters = []Route{
	{
		URI:                    "/reports/GetRentedBooks",
		Method:                 http.MethodGet,
		Function:               handlers.GetRentedBooks,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/reports/GetNumberOfOverdueBooks",
		Method:                 http.MethodGet,
		Function:               handlers.GetNumberOfOverdueBooks,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/reports/GetNumberOfBooksRentsByClient/{clientName}",
		Method:                 http.MethodGet,
		Function:               handlers.GetNumberOfBooksRentsByClient,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/reports/GetReturnedBooks",
		Method:                 http.MethodGet,
		Function:               handlers.GetReturnedBooks,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/reports/GetRentsByClient/{clientName}",
		Method:                 http.MethodGet,
		Function:               handlers.GetNumberRentsByClient,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/reports/GetMostRentedBook",
		Method:                 http.MethodGet,
		Function:               handlers.GetMostRentedBook,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/reports/GetAllReports",
		Method:                 http.MethodGet,
		Function:               handlers.GetAllReports,
		RequiresAuthentication: true,
	},
}

func init() {
	Routes = append(Routes, reportsRouters...)
}
