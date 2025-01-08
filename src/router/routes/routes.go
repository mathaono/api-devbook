package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Estrutura que representa todas as rotas da API
type Rota struct {
	URI            string
	Method         string
	Func           func(http.ResponseWriter, *http.Request)
	Authentication bool
}

// Coloca todas as rotas dentro do router (roteador)
func Config(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
