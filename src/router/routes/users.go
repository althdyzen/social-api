package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Rota{
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
		NeedAuth: false,
	},
	{
		URI:      "/users",
		Method:   http.MethodGet,
		Function: controllers.GetUsers,
		NeedAuth: true,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodGet,
		Function: controllers.GetUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{id}/follow",
		Method:   http.MethodPost,
		Function: controllers.FollowUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{id}/unfollow",
		Method:   http.MethodPost,
		Function: controllers.UnfollowUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{id}/followers",
		Method:   http.MethodGet,
		Function: controllers.Followers,
		NeedAuth: true,
	},
	{
		URI:      "/users/{id}/following",
		Method:   http.MethodGet,
		Function: controllers.Following,
		NeedAuth: true,
	},
	{
		URI:      "/users/{id}/update-password",
		Method:   http.MethodPost,
		Function: controllers.UpdatePassword,
		NeedAuth: true,
	},
	{
		URI:      "/isauthenticated",
		Method:   http.MethodGet,
		Function: controllers.IsAuthenticated,
		NeedAuth: true,
	},
}
