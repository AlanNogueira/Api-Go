package handlers

import (
	"Api-Go/pkg/entities"
	"Api-Go/pkg/entities/utils"
	"Api-Go/pkg/repositories"
	"Api-Go/pkg/responses"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateClient(w http.ResponseWriter, r *http.Request) {
	var client entities.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := client.Prepare(); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	clientRepository := repositories.CreateNewClientRepository()
	response, err := clientRepository.Create(client)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func GetClients(w http.ResponseWriter, r *http.Request) {
	clientRepository := repositories.CreateNewClientRepository()
	options := options.Find()
	_, _ = utils.Pagination(r, options)
	response, err := clientRepository.GetClients(options)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func GetClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taxNumber := params["taxNumber"]

	clientRepository := repositories.CreateNewClientRepository()
	response, err := clientRepository.GetClient(taxNumber)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["clientId"]

	var client entities.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	clientRepository := repositories.CreateNewClientRepository()
	response, err := clientRepository.UpdateClient(id, client)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func RemoveClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["clientId"]

	clientRepository := repositories.CreateNewClientRepository()
	response, err := clientRepository.RemoveClient(id)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}
