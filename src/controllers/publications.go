package controllers

import (
	"api/src/authentication"
	"api/src/config/repositories"
	"api/src/database"
	"api/src/models"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Permite que o usuário crie uma nova publicação
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		log.Printf("[ERROR] - User unauthorized: %v", err)
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication
	if err = json.Unmarshal(requestBody, &publication); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	publication.UserID = userID

	if err = publication.Prepare(); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationRepository(db)
	publication.ID, err = repository.Create(publication)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, publication)
}

// Permite que o usuário pesquise uma publicação
func SearchPublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
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

	repository := repositories.NewPublicationRepository(db)
	pubs, err := repository.SearchPublications(userID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, pubs)
}

// Permite que o usuário busque uma publicação pelo ID
func SearchPublicationByID(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	pubID, err := strconv.ParseUint(parameters["publicationID"], 10, 64)
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

	repository := repositories.NewPublicationRepository(db)
	pub, err := repository.SearchPublicationByID(pubID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, pub)
}

// Permite que o usuário atualize uma publicação
func UpdatePublication(w http.ResponseWriter, r *http.Request) {

}

// Permite que o usuário exclua uma publicação
func DeletePublication(w http.ResponseWriter, r *http.Request) {

}
