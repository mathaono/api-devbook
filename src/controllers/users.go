package controllers

import (
	"api/src/config/repositories"
	"api/src/database"
	"api/src/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Lida com a função HTTP (POST) de CRIAR um usuário e manda para o pacote repositories CRIAR o usuário na tabela
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Erro ao capturar dados do corpo da request: ", err)
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal("Erro ao executar Unmarshal: ", err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados: ", err)
	}

	repository := repositories.NewUsersRepository(db)
	userId, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("ID inserido: %d", userId)))
}

// Lida com a função HTTP (GET) de BUSCAR usuários e manda para o pacote repositories BUSCAR usuários na tabela
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuários!"))
}

// Lida com a função HTTP (GET) de BUSCAR um usuário PELO ID e manda para o pacote repositories BUSCAR o usuário PELO ID na tabela
func SearchUserByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário por ID!"))
}

// Lida com a função HTTP (PUT) de ATUALIZAR um usuário e manda para o pacote repositories ATUALIZAR o usuário na tabela
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário!"))
}

// Lida com a função HTTP (DELETE) de DELETAR um usuário e manda para o pacote repositories DELETAR o usuário na tabela
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Excluindo usuário!"))
}
