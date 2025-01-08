package middlewares

import (
	"fmt"
	"log"
	"net/http"
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
		fmt.Println("Autenticando")
		next(w, r)
	}
}
