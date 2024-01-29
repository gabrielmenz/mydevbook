package routes

import (
	"goapi/src/controllers"
	"net/http"
)

var RoutesUsers = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.SearchUsers,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodGet,
		Function:              controllers.SearchUser,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users/{userId}/update",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUser,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users/{userId}/delete",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		RequestAuthentication: false,
	},
}
