package controllers

import (
	"api/src/authentication"
	"api/src/config/repositories"
	"api/src/database"
	"api/src/models"
	"api/src/responses"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Lida com a função HTTP (POST) de CRIAR um usuário e manda para o pacote repositories CRIAR o usuário na tabela
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Printf("[ERROR] - Failed Unmarshal JSON: %v", err)
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	log.Printf("[RECEIVED] - Creating user: %v", user)

	if err = user.Prepare("create"); err != nil {
		log.Printf("[ERROR] - Failed to prepare user: %v", err)
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	userId, err := repository.Create(user)
	if err != nil {
		log.Printf("[ERROR] - Failed to create user: %v", err)
		responses.Erro(w, http.StatusInternalServerError, err)
	}

	user.ID = int(userId)

	log.Printf("[SUCCESS] - User created: %v", user)

	responses.JSON(w, http.StatusCreated, user)
}

// Lida com a função HTTP (GET) de BUSCAR usuários e manda para o pacote repositories BUSCAR usuários na tabela
func SearchUser(w http.ResponseWriter, r *http.Request) {
	userName := strings.ToLower(r.URL.Query().Get("user"))
	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	users, err := repository.Search(userName)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// Lida com a função HTTP (GET) de BUSCAR um usuário PELO ID e manda para o pacote repositories BUSCAR o usuário PELO ID na tabela
func SearchUserByID(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["ID"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user, err := repository.SearchByID(userID)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// Lida com a função HTTP (PUT) de ATUALIZAR um usuário e manda para o pacote repositories ATUALIZAR o usuário na tabela
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["ID"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	userTokenID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	fmt.Println("user ID Token: ", userTokenID)
	if userID != userTokenID {
		log.Println("[ERROR] - UserID and token are different")
		responses.Erro(w, http.StatusForbidden, errors.New("UserID and token are different"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("update"); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	err = repository.Update(userID, user)

	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Lida com a função HTTP (DELETE) de DELETAR um usuário e manda para o pacote repositories DELETAR o usuário na tabela
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["ID"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	userTokenID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userTokenID {
		log.Println("[ERROR] - UserID and token are different")
		responses.Erro(w, http.StatusForbidden, errors.New("UserID and token are different"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	err = repository.Delete(userID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
