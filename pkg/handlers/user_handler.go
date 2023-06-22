package handlers

import (
	"Api-Go/pkg/entities"
	"Api-Go/pkg/repositories"
	"Api-Go/pkg/responses"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := user.Prepare(); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	userRepository := repositories.CreateNewUserRepository()
	response, err := userRepository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userRepository := repositories.CreateNewUserRepository()
	response, err := userRepository.GetUsers()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	userRepository := repositories.CreateNewUserRepository()
	response, err := userRepository.GetUser(name)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["userId"]

	var user entities.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	userRepository := repositories.CreateNewUserRepository()
	response, err := userRepository.UpdateUser(id, user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["userId"]

	userRepository := repositories.CreateNewUserRepository()
	response, err := userRepository.RemoveUser(id)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}
