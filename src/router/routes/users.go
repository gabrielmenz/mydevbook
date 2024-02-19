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
		RequestAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodGet,
		Function:              controllers.SearchUser,
		RequestAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUser,
		RequestAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		RequestAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/follow",
		Method:                http.MethodPost,
		Function:              controllers.FollowUser,
		RequestAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/unfollow",
		Method:                http.MethodPost,
		Function:              controllers.UnfollowUser,
		RequestAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/followers",
		Method:                http.MethodGet,
		Function:              controllers.SearchFollowers,
		RequestAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/following",
		Method:                http.MethodGet,
		Function:              controllers.SearchFollowing,
		RequestAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/update-pw",
		Method:                http.MethodPost,
		Function:              controllers.UpdatePw,
		RequestAuthentication: true,
	},
}
