package routes

import (
	"Api-Go/pkg/handlers"
	"net/http"
)

var bookRoutes = []Route{
	{
		URI:                    "/books",
		Method:                 http.MethodPost,
		Function:               handlers.CreateBook,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/books",
		Method:                 http.MethodGet,
		Function:               handlers.GetBooks,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/books/{name}",
		Method:                 http.MethodGet,
		Function:               handlers.GetBook,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/books/{bookId}",
		Method:                 http.MethodPut,
		Function:               handlers.UpdateBook,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/books/{bookId}",
		Method:                 http.MethodDelete,
		Function:               handlers.RemoveBook,
		RequiresAuthentication: false,
	},
}

func init() {
	Routes = append(Routes, bookRoutes...)
}
