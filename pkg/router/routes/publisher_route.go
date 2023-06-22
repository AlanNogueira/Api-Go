package routes

import (
	"Api-Go/pkg/handlers"
	"net/http"
)

var publisherRoutes = []Route{
	{
		URI:                    "/publishers",
		Method:                 http.MethodPost,
		Function:               handlers.CreatePublisher,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publishers",
		Method:                 http.MethodGet,
		Function:               handlers.GetPublishers,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publishers/{name}",
		Method:                 http.MethodGet,
		Function:               handlers.GetPublisher,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publishers/{publisherId}",
		Method:                 http.MethodPut,
		Function:               handlers.UpdatePublisher,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publishers/{publisherId}",
		Method:                 http.MethodDelete,
		Function:               handlers.RemovePublisher,
		RequiresAuthentication: true,
	},
}

func init() {
	Routes = append(Routes, publisherRoutes...)
}
