package controllers

import (
	"api/security"
	"api/src/authentication"
	"api/src/config/repositories"
	"api/src/database"
	"api/src/models"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

// Autenticação de usuário na API
func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	err = json.Unmarshal(requestBody, &user)
	if err != nil {
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
	userDB, err := repository.SearchByEmail(user.Email)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	err = security.PasswordValidate(userDB.Password, user.Password)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(uint64(userDB.ID))
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))
}
