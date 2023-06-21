package handlers

import (
	"Api-Go/pkg/entities"
	"Api-Go/pkg/repositories"
	"Api-Go/pkg/responses"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRent(w http.ResponseWriter, r *http.Request) {
	var rent entities.Rent
	err := json.NewDecoder(r.Body).Decode(&rent)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := rent.Prepare(); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	rentRepository := repositories.CreateNewRentRepository()
	response, err := rentRepository.CreateRent(rent)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func GetNotDeliveredRents(w http.ResponseWriter, r *http.Request) {
	rentRepository := repositories.CreateNewRentRepository()
	reponse, err := rentRepository.GetNotDeliveredRents()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, reponse)
}

func GetRent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	filters := map[string]string{
		"userName":      params["userName"],
		"bookAuthor":    r.URL.Query().Get("bookAuthor"),
		"bookPublisher": r.URL.Query().Get("bookPublisher"),
		"bookName":      r.URL.Query().Get("bookName"),
		"rentStatus":    r.URL.Query().Get("rentStatus"),
	}

	rentRepository := repositories.CreateNewRentRepository()
	response, err := rentRepository.GetRent(filters)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, response)
}

func FinalizeRent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["rentId"]

	rentRepository := repositories.CreateNewRentRepository()
	response, err := rentRepository.FinalizeRent(id)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, response)
}
