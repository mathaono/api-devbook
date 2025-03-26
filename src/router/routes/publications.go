package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesPublications = []Rota{
	{
		URI:            "/publications",
		Method:         http.MethodPost,
		Func:           controllers.CreatePublication,
		Authentication: true,
	},
	{
		URI:            "/publications",
		Method:         http.MethodGet,
		Func:           controllers.SearchPublication,
		Authentication: true,
	},
	{
		URI:            "/publications/{publicationID}",
		Method:         http.MethodGet,
		Func:           controllers.SearchPublicationByID,
		Authentication: true,
	},
	{
		URI:            "/publications/{publicationID}",
		Method:         http.MethodPut,
		Func:           controllers.UpdatePublication,
		Authentication: true,
	},
	{
		URI:            "/publications/{publicationID}",
		Method:         http.MethodDelete,
		Func:           controllers.DeletePublication,
		Authentication: true,
	},
	{
		URI:            "/users/{userID}/publications",
		Method:         http.MethodGet,
		Func:           controllers.SearchPublicationByUser,
		Authentication: true,
	},
	{
		URI:            "/publications/{publicationID}/like",
		Method:         http.MethodPost,
		Func:           controllers.LikePublication,
		Authentication: true,
	},
	{
		URI:            "/publications/{publicationID}/dislike",
		Method:         http.MethodPost,
		Func:           controllers.DislikePublication,
		Authentication: true,
	},
}
