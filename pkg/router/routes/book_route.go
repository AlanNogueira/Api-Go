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
		RequiresAuthentication: true,
	},
	{
		URI:                    "/books",
		Method:                 http.MethodGet,
		Function:               handlers.GetBooks,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/books/{name}",
		Method:                 http.MethodGet,
		Function:               handlers.GetBook,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/books/{bookId}",
		Method:                 http.MethodPut,
		Function:               handlers.UpdateBook,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/books/{bookId}",
		Method:                 http.MethodDelete,
		Function:               handlers.RemoveBook,
		RequiresAuthentication: true,
	},
}

func init() {
	Routes = append(Routes, bookRoutes...)
}
