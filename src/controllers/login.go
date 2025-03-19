package controllers

import (
	"api/src/authentication"
	"api/src/config/repositories"
	"api/src/database"
	"api/src/models"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Autenticação de usuário na API
func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("[LOGIN] - Receive login request")
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("[ERROR] - Failed to read request body: %v", err)
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	err = json.Unmarshal(requestBody, &user)
	if err != nil {
		log.Printf("[ERROR] - Failed to unmarshal JSON: %v", err)
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	log.Printf("[LOGIN] - Trying to authenticate user: %s", user.Email)

	db, err := database.Connect()
	if err != nil {
		log.Printf("[ERROR DB] - Failed to connect database: %v", err)
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	log.Println("[PROCESS] - Connect database")

	repository := repositories.NewUsersRepository(db)
	userDB, err := repository.SearchByEmail(user.Email)
	if err != nil {
		log.Printf("[ERROR DB] - User not found: %v", err)
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	err = security.PasswordValidate(userDB.Password, user.Password)
	if err != nil {
		log.Printf("[ERROR] - Invalid password for user: %s", user.Email)
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(uint64(userDB.ID))
	if err != nil {
		log.Printf("[ERROR] - Failed to token generate for user: %s - %v", user.Email, err)
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	log.Printf("[SUCCESS] User authenticate: %s", user.Email)
	w.Write([]byte(token))
}
