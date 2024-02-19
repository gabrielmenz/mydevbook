package middlewares

import (
	"goapi/src/authentication"
	"goapi/src/responses"
	"log"
	"net/http"
)

// Middlewares role is to implement a function to all routes, instead of entering the function inside EACH route

// Logger writes request info on the terminal
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

// Verifies if the user is authenticated
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			responses.Error(w, http.StatusUnauthorized, err)
			return
		}
		nextFunction(w, r)
	}
}
