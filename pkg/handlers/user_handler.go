package handlers

import (
	"Api-Go/pkg/configuration"
	"Api-Go/pkg/entities"
	"Api-Go/pkg/repositories"
	"Api-Go/pkg/responses"
	"encoding/json"
	"fmt"
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

	collection, err := configuration.DbConnect("users")
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.CreateNewUserRepository(collection)
	response, err := userRepository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	if err := enc.Encode(response); err != nil {
		fmt.Println(err.Error())
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	collection, err := configuration.DbConnect("users")
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.CreateNewUserRepository(collection)
	response, err := userRepository.GetUsers()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	if err := enc.Encode(response); err != nil {
		fmt.Println(err.Error())
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	collection, err := configuration.DbConnect("users")
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.CreateNewUserRepository(collection)
	response, err := userRepository.GetUser(name)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	if err := enc.Encode(response); err != nil {
		fmt.Println(err.Error())
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

	collection, err := configuration.DbConnect("users")
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.CreateNewUserRepository(collection)
	response, err := userRepository.UpdateUser(id, user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	if err := enc.Encode(response); err != nil {
		fmt.Println(err.Error())
	}
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["userId"]

	collection, err := configuration.DbConnect("users")
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.CreateNewUserRepository(collection)
	response, err := userRepository.RemoveUser(id)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	if err := enc.Encode(response); err != nil {
		fmt.Println(err.Error())
	}
}
