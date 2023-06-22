package handlers

import (
	"Api-Go/pkg/entities"
	"Api-Go/pkg/repositories"
	"Api-Go/pkg/responses"
	"encoding/json"
	"net/http"
)

func CreateRentalCompany(w http.ResponseWriter, r *http.Request) {
	var rentalCompany entities.RentalCompany
	if err := json.NewDecoder(r.Body).Decode(&rentalCompany); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := rentalCompany.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	rentalCompanyRepository := repositories.CreateNewRentalCompanyRepository()
	response, err := rentalCompanyRepository.CreateRentalCompany(rentalCompany)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var rentalCompany entities.RentalCompany
	if err := json.NewDecoder(r.Body).Decode(&rentalCompany); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	rentalCompanyRepository := repositories.CreateNewRentalCompanyRepository()
	response, err := rentalCompanyRepository.AuthenticateUser(rentalCompany)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	responses.JSON(w, http.StatusOK, response)
}
