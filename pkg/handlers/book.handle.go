package handlers

import (
	"Api-Go/pkg/configuration"
	"Api-Go/pkg/models"
	"Api-Go/pkg/repositories"
	"Api-Go/pkg/responses"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	collection, err := configuration.DbConnect("books")
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	bookRepository := repositories.CreateNewBooksRepository(collection)
	response, err := bookRepository.Create(book)
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

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["bookId"]

	collection, err := configuration.DbConnect("books")
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	bookRepository := repositories.CreateNewBooksRepository(collection)
	response, err := bookRepository.GetBook(id)
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

func GetBooks(w http.ResponseWriter, r *http.Request) {
	collection, err := configuration.DbConnect("books")
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	bookRepository := repositories.CreateNewBooksRepository(collection)
	response, err := bookRepository.GetBooks()
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

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["bookId"]

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	collection, err := configuration.DbConnect("books")
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	bookRepository := repositories.CreateNewBooksRepository(collection)
	response, err := bookRepository.UpdateBook(id, book)
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

func RemoveBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["bookId"]

	collection, err := configuration.DbConnect("books")
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	bookRepository := repositories.CreateNewBooksRepository(collection)
	response, err := bookRepository.RemoveBook(id)
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
