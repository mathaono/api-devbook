package controllers

import "net/http"

// Insere um usuário no banco
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário!"))
}

// Busca todos os usuários no banco
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuários!"))
}

// Busca usuário no banco pelo ID
func SearchUserByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário por ID!"))
}

// Altera as infos de um usuário no banco
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário!"))
}

// Exclui um usuário do banco
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Excluindo usuário!"))
}
