package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
	"log"
	"net/http"
	"time"
)

// Escreve no terminal as informações da request
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Verifica se a requisição do usuário está autenticada
func Autenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := authentication.ValidateToken(r)
		if err != nil {
			responses.Erro(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}

func LogRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[REQUEST] - METHOD: %s - PATH: %s - IP: %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("[COMPLETED] - METHOD: %s - PATH: %s - Time: %s", r.Method, r.URL.Path, time.Since(start))
	})
}
