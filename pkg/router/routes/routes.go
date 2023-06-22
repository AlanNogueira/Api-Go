package routes

import (
	"Api-Go/pkg/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

var Routes = []Route{}

type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

func Configure(r *mux.Router) *mux.Router {
	for _, route := range Routes {

		if route.RequiresAuthentication {
			r.HandleFunc(route.URI, middlewares.Authenticate(route.Function)).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, route.Function).Methods(route.Method)
		}

	}

	return r
}
