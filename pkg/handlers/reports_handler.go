package handlers

import (
	"Api-Go/pkg/repositories"
	"Api-Go/pkg/responses"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRentedBooks(w http.ResponseWriter, r *http.Request) {
	rentsRepository := repositories.CreateNewRentRepository()
	response, err := rentsRepository.GetRentedBooks()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, response)
}

func GetNumberOfOverdueBooks(w http.ResponseWriter, r *http.Request) {
	rentsRepository := repositories.CreateNewRentRepository()
	response, err := rentsRepository.GetNumberOfOverdueBooks()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, response)
}

func GetNumberOfBooksRentsByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userName := params["userName"]
	rentsRepository := repositories.CreateNewRentRepository()
	response, err := rentsRepository.GetNumberOfBooksRentsByUser(userName)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, response)
}

func GetReturnedBooks(w http.ResponseWriter, r *http.Request) {
	rentsRepository := repositories.CreateNewRentRepository()
	response, err := rentsRepository.GetReturnedBooks()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, response)
}

func GetNumberRentsByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userName := params["userName"]

	rentRepository := repositories.CreateNewRentRepository()
	reponse, err := rentRepository.GetNumberRentsByUser(userName)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}

	responses.JSON(w, http.StatusOK, reponse)
}

func GetMostRentedBook(w http.ResponseWriter, r *http.Request) {
	bookRepository := repositories.CreateNewBooksRepository()
	response, err := bookRepository.GetMostRentedBook()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, response)
}
