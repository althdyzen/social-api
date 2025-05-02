package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesPosts = []Rota{
	{
		URI:      "/posts",
		Method:   http.MethodPost,
		Function: controllers.CreatePost,
		NeedAuth: true,
	},
	{
		URI:      "/posts",
		Method:   http.MethodGet,
		Function: controllers.GetPosts,
		NeedAuth: true,
	},
	{
		URI:      "/posts/me",
		Method:   http.MethodGet,
		Function: controllers.GetPostsMe,
		NeedAuth: true,
	},
	{
		URI:      "/posts/{id}",
		Method:   http.MethodGet,
		Function: controllers.GetPost,
		NeedAuth: true,
	},
	{
		URI:      "/posts/{id}",
		Method:   http.MethodPut,
		Function: controllers.UpdatePost,
		NeedAuth: true,
	},
	{
		URI:      "/posts/{id}",
		Method:   http.MethodDelete,
		Function: controllers.DeletePost,
		NeedAuth: true,
	},
	{
		URI:      "/users/{id}/posts",
		Method:   http.MethodGet,
		Function: controllers.GetPostByUser,
		NeedAuth: true,
	},
	{
		URI:      "/posts/{id}/like",
		Method:   http.MethodPost,
		Function: controllers.LikePost,
		NeedAuth: true,
	},
	{
		URI:      "/posts/{id}/dislike",
		Method:   http.MethodPost,
		Function: controllers.DislikePost,
		NeedAuth: true,
	},
}
