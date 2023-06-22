package handlers

import (
	"Api-Go/pkg/entities"
	"Api-Go/pkg/entities/utils"
	"Api-Go/pkg/repositories"
	"Api-Go/pkg/responses"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book entities.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := book.Prepare(); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	bookRepository := repositories.CreateNewBooksRepository()
	response, err := bookRepository.Create(book)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	author := r.URL.Query().Get("author")
	publisher := r.URL.Query().Get("publisher")
	releaseDate := r.URL.Query().Get("releaseDate")
	time, _ := time.Parse("02-01-2006", releaseDate)

	book := entities.Book{
		Name:        name,
		Author:      author,
		Publisher:   publisher,
		ReleaseDate: utils.CustomTime{Time: time},
	}

	bookRepository := repositories.CreateNewBooksRepository()
	response, err := bookRepository.GetBook(book)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	bookRepository := repositories.CreateNewBooksRepository()
	response, err := bookRepository.GetBooks()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["bookId"]

	var book entities.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	bookRepository := repositories.CreateNewBooksRepository()
	response, err := bookRepository.UpdateBook(id, book)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func RemoveBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["bookId"]

	bookRepository := repositories.CreateNewBooksRepository()
	response, err := bookRepository.RemoveBook(id)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}
