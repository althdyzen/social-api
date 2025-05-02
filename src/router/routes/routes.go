package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI      string
	Method   string
	Function func(w http.ResponseWriter, r *http.Request)
	NeedAuth bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, routesPosts...)

	for _, route := range routes {

		if route.NeedAuth {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method, "OPTIONS")
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method, "OPTIONS")
		}
	}

	return r
}
