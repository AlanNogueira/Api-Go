package handlers

import (
	"Api-Go/pkg/entities"
	"Api-Go/pkg/repositories"
	"Api-Go/pkg/responses"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreatePublisher(w http.ResponseWriter, r *http.Request) {
	var publisher entities.Publisher
	if err := json.NewDecoder(r.Body).Decode(&publisher); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := publisher.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	publisherRepository := repositories.CreateNewPublishersRepository()
	response, err := publisherRepository.CreatePublisher(publisher)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func GetPublishers(w http.ResponseWriter, r *http.Request) {
	publisherRepository := repositories.CreateNewPublishersRepository()
	response, err := publisherRepository.GetPublishers()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func GetPublisher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	publisherRepository := repositories.CreateNewPublishersRepository()
	response, err := publisherRepository.GetPublisher(name)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func UpdatePublisher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Id := params["publisherId"]

	var publisher entities.Publisher
	if err := json.NewDecoder(r.Body).Decode(&publisher); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	publisherRepository := repositories.CreateNewPublishersRepository()
	response, err := publisherRepository.UpdatePublisher(Id, publisher)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func RemovePublisher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Id := params["publisherId"]

	publisherRepository := repositories.CreateNewPublishersRepository()
	response, err := publisherRepository.RemovePublisher(Id)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}
