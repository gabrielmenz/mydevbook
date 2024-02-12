package routes

import (
	"goapi/src/controllers"
	"net/http"
)

var routeLogin = Route{
	URI:                   "/login",
	Method:                http.MethodPost,
	Function:              controllers.Login,
	RequestAuthentication: false,
}
