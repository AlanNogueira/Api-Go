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

	publisherRepository, err := repositories.CreateNewPublishersRepository()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	response, err := publisherRepository.CreatePublisher(publisher)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	if err := enc.Encode(response); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
}

func GetPublishers(w http.ResponseWriter, r *http.Request) {
	publisherRepository, err := repositories.CreateNewPublishersRepository()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	response, err := publisherRepository.GetPublishers()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	if err := enc.Encode(response); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
}

func GetPublisher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	publisherRepository, err := repositories.CreateNewPublishersRepository()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	response, err := publisherRepository.GetPublisher(name)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	if err := enc.Encode(response); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
}

func UpdatePublisher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Id := params["publisherId"]

	var publisher entities.Publisher
	if err := json.NewDecoder(r.Body).Decode(&publisher); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	publisherRepository, err := repositories.CreateNewPublishersRepository()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	response, err := publisherRepository.UpdatePublisher(Id, publisher)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	if err := enc.Encode(response); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
}

func RemovePublisher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Id := params["publisherId"]

	publisherRepository, err := repositories.CreateNewPublishersRepository()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	response, err := publisherRepository.RemovePublisher(Id)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	if err := enc.Encode(response); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
}
