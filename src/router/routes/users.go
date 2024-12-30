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
		Authentication: false,
	},
	{
		URI:            "/users/{ID}",
		Method:         http.MethodGet,
		Func:           controllers.SearchUserByID,
		Authentication: false,
	},
	{
		URI:            "/users/{ID}",
		Method:         http.MethodPut,
		Func:           controllers.UpdateUser,
		Authentication: false,
	},
	{
		URI:            "/users/{ID}",
		Method:         http.MethodDelete,
		Func:           controllers.DeleteUser,
		Authentication: false,
	},
}
