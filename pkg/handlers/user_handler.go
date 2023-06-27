package handlers

import (
	"Api-Go/pkg/entities"
	"Api-Go/pkg/repositories"
	"Api-Go/pkg/responses"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := user.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userRepository := repositories.CreateNewUserRepository()
	response, err := userRepository.CreateUser(user)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	userRepository := repositories.CreateNewUserRepository()
	response, err := userRepository.AuthenticateUser(user)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	responses.JSON(w, http.StatusOK, response)
}
