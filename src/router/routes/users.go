package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Rota{
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Func:           controllers.CreateUser,
		Authentication: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Func:           controllers.SearchUser,
		Authentication: true,
	},
	{
		URI:            "/users/{ID}",
		Method:         http.MethodGet,
		Func:           controllers.SearchUserByID,
		Authentication: true,
	},
	{
		URI:            "/users/{ID}",
		Method:         http.MethodPut,
		Func:           controllers.UpdateUser,
		Authentication: true,
	},
	{
		URI:            "/users/{ID}",
		Method:         http.MethodDelete,
		Func:           controllers.DeleteUser,
		Authentication: true,
	},
	{
		URI:            "/users/{ID}/follow",
		Method:         http.MethodPost,
		Func:           controllers.FollowUser,
		Authentication: true,
	},
	{
		URI:            "/users/{ID}/unfollow",
		Method:         http.MethodPost,
		Func:           controllers.UnfollowUser,
		Authentication: true,
	},
	{
		URI:            "/users/{ID}/followers",
		Method:         http.MethodGet,
		Func:           controllers.SearchFollowers,
		Authentication: true,
	},
}
