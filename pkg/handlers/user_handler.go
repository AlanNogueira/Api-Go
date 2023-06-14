package handlers

import (
	"Api-Go/pkg/configuration"
	"Api-Go/pkg/models"
	"Api-Go/pkg/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var usuario models.User
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		log.Fatal(err)
	}

	collection, err := configuration.DbConnect("users")
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repositories.CreateNewUserRepository(collection)
	response, err := userRepository.Create(usuario)
	if err != nil {
		response = map[string]interface{}{"error": err.Error()}
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
		log.Fatal(err)
	}

	userRepository := repositories.CreateNewUserRepository(collection)
	response, err := userRepository.GetUsers()
	if err != nil {
		response = map[string]interface{}{"error": err.Error()}
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	if err := enc.Encode(response); err != nil {
		fmt.Println(err.Error())
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["userId"]

	collection, err := configuration.DbConnect("users")
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repositories.CreateNewUserRepository(collection)
	response, err := userRepository.GetUser(id)
	if err != nil {
		response = map[string]interface{}{"error": err.Error()}
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

	var usuario models.User
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		log.Fatal(err)
	}

	collection, err := configuration.DbConnect("users")
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repositories.CreateNewUserRepository(collection)
	response, err := userRepository.UpdateUser(id, usuario)
	if err != nil {
		response = map[string]interface{}{"error": err.Error()}
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
		log.Fatal(err)
	}

	userRepository := repositories.CreateNewUserRepository(collection)
	response, err := userRepository.RemoveUser(id)
	if err != nil {
		response = map[string]interface{}{"error": err.Error()}
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	if err := enc.Encode(response); err != nil {
		fmt.Println(err.Error())
	}
}
