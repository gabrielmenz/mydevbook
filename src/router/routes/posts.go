package routes

import (
	"goapi/src/controllers"
	"net/http"
)

var PostsRoutes = []Route{
	{
		URI:                   "/posts",
		Method:                http.MethodPost,
		Function:              controllers.CreatePost,
		RequestAuthentication: true,
	},
	{
		URI:                   "/posts",
		Method:                http.MethodGet,
		Function:              controllers.SearchPosts,
		RequestAuthentication: true,
	},
	{
		URI:                   "/posts/{postId}",
		Method:                http.MethodGet,
		Function:              controllers.SearchPost,
		RequestAuthentication: true,
	},
	{
		URI:                   "/posts/{postId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdatePost,
		RequestAuthentication: true,
	},
	{
		URI:                   "/posts/{postId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeletePost,
		RequestAuthentication: true,
	},
	{
		URI:                   "/usuarios/{usuarioId}/posts",
		Method:                http.MethodGet,
		Function:              controllers.SearchPostsByUser,
		RequestAuthentication: true,
	},
	{
		URI:                   "/posts/{postId}/like",
		Method:                http.MethodPost,
		Function:              controllers.LikePost,
		RequestAuthentication: true,
	},
	{
		URI:                   "/posts/{postId}/dislike",
		Method:                http.MethodPost,
		Function:              controllers.DislikePost,
		RequestAuthentication: true,
	},
}
