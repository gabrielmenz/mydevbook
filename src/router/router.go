package router

import (
	"goapi/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate returns router with configured routes.
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configurate(r)

}
