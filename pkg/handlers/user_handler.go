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
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userRepository, err := repositories.CreateNewUserRepository()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	response, err := userRepository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	if err := enc.Encode(response); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userRepository, err := repositories.CreateNewUserRepository()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	response, err := userRepository.GetUsers()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	if err := enc.Encode(response); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	userRepository, err := repositories.CreateNewUserRepository()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	response, err := userRepository.GetUser(name)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	if err := enc.Encode(response); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["userId"]

	var user entities.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userRepository, err := repositories.CreateNewUserRepository()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	response, err := userRepository.UpdateUser(id, user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	if err := enc.Encode(response); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["userId"]

	userRepository, err := repositories.CreateNewUserRepository()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	response, err := userRepository.RemoveUser(id)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	if err := enc.Encode(response); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
}
