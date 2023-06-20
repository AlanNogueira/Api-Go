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
		RequiresAuthentication: false,
	},
	{
		URI:                    "/publishers",
		Method:                 http.MethodGet,
		Function:               handlers.GetPublishers,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/publishers/{name}",
		Method:                 http.MethodGet,
		Function:               handlers.GetPublisher,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/publishers/{publisherId}",
		Method:                 http.MethodPut,
		Function:               handlers.UpdatePublisher,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/publishers/{publisherId}",
		Method:                 http.MethodDelete,
		Function:               handlers.RemovePublisher,
		RequiresAuthentication: false,
	},
}

func init() {
	Routes = append(Routes, publisherRoutes...)
}
