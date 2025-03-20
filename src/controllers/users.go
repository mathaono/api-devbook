package controllers

import (
	"api/src/authentication"
	"api/src/config/repositories"
	"api/src/database"
	"api/src/models"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
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

// Permite que um usuário siga outro
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		log.Printf("[ERROR] - Unauthorized to following: %v", err)
		responses.Erro(w, http.StatusUnauthorized, err)
	}

	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["ID"], 10, 64)
	if err != nil {
		log.Printf("[ERROR] - userID is required: %v", err)
		responses.Erro(w, http.StatusBadRequest, err)
	}

	if followerID == userID {
		responses.Erro(w, http.StatusForbidden, errors.New("[ERROR] - followerID and userID are the same"))
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if err = repository.Follow(userID, followerID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Permite deixar de seguir um usuário
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		log.Printf("[ERROR] - Unauthorized to unfollowing: %v", err)
		responses.Erro(w, http.StatusUnauthorized, err)
	}

	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["ID"], 10, 64)
	if err != nil {
		log.Printf("[ERROR] - userID is required: %v", err)
		responses.Erro(w, http.StatusBadRequest, err)
	}

	if followerID == userID {
		responses.Erro(w, http.StatusForbidden, errors.New("[ERROR] - followerID and userID are the same"))
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if err = repository.Unfollow(userID, followerID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Permite fazer uma busca de seguidores de um usuário
func SearchFollowers(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["ID"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	followers, err := repository.SearchFollowers(userID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

// Trás todos os usuários que o usuário da request está seguindo
func SearchFollowing(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["ID"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	users, err := repository.SearchFollowing(userID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// Permite que o usuário altere sua senha
func ResetPass(w http.ResponseWriter, r *http.Request) {
	userTokenID, err := authentication.ExtractUserID(r)
	if err != nil {
		log.Printf("[ERROR] - Unauthorized to unfollowing: %v", err)
		responses.Erro(w, http.StatusUnauthorized, err)
	}

	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["ID"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if userTokenID != userID {
		responses.Erro(w, http.StatusForbidden, errors.New("[ERROR] - userTokenID and userID are different"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)

	var pass models.Pass
	if err = json.Unmarshal(requestBody, &pass); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	passOnDB, err := repository.SearchPass(userID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.PasswordValidate(passOnDB, pass.CurrentPass); err != nil {
		responses.Erro(w, http.StatusUnauthorized, errors.New("[ERROR] - current password and password on DB are different"))
		return
	}

	passWithHash, err := security.Hash(pass.NewPass)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.UpdatePass(userID, string(passWithHash)); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
