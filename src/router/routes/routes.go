package routes

import (
	"goapi/src/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequestAuthentication bool
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("\n %s 5s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// sends all routes to the router
func Configurate(r *mux.Router) *mux.Router {
	routes := RoutesUsers
	routes = append(routes, routeLogin)
	routes = append(routes, PostsRoutes...)

	for _, route := range routes {
		if route.RequestAuthentication {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}
	return r

}
