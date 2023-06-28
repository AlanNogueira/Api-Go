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
	"go.mongodb.org/mongo-driver/mongo/options"
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
	releaseDate := r.URL.Query().Get("releaseDate")
	if releaseDate != "" {
		time, _ := time.Parse("02-01-2006", releaseDate)
		releaseDate = time.Format("02-01-2006")
	}
	filters := map[string]string{
		"name":        params["name"],
		"author":      r.URL.Query().Get("author"),
		"publisher":   r.URL.Query().Get("publisher"),
		"releaseDate": releaseDate,
	}

	bookRepository := repositories.CreateNewBooksRepository()
	response, err := bookRepository.GetBook(filters)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	bookRepository := repositories.CreateNewBooksRepository()
	options := options.Find()
	_, _ = utils.Pagination(r, options)
	response, err := bookRepository.GetBooks(options)
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
