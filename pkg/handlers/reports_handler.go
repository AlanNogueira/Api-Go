package handlers

import (
	"Api-Go/pkg/repositories"
	"Api-Go/pkg/responses"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetAllReports(w http.ResponseWriter, r *http.Request) {
	rentsRepository := repositories.CreateNewRentRepository()
	rentedBooks, err := rentsRepository.GetRentedBooks()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	overdueBooks, err := rentsRepository.GetNumberOfOverdueBooks()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	var rentedBooksByUser []map[string]interface{}
	var numRentsByUser []map[string]interface{}
	userRepository := repositories.CreateNewUserRepository()
	options := options.Find()
	users, err := userRepository.GetUsers(options)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	for i := 0; i < len(users); i++ {
		response, err := rentsRepository.GetNumberOfBooksRentsByUser(users[i].Name)
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}
		rentedBooksByUser = append(rentedBooksByUser, response)

		reponse, err := rentsRepository.GetNumberRentsByUser(users[i].Name)
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}
		numRentsByUser = append(numRentsByUser, reponse)
	}

	returnedBooks, err := rentsRepository.GetReturnedBooks()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	bookRepository := repositories.CreateNewBooksRepository()
	mostRentedBook, err := bookRepository.GetMostRentedBook()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	res := map[string]interface{}{
		"data": map[string]interface{}{
			"NumberOfRentedBooks":      rentedBooks["NumberOfRentedBooks"],
			"NumberOfOverdueBooks":     overdueBooks["NumberOfOverdueBooks"],
			"NumberOfBooksRentsByUser": rentedBooksByUser,
			"ReturnedBooks":            returnedBooks,
			"NumberRentsByUser":        numRentsByUser,
			"MostRentedBook":           mostRentedBook["mostRentedBook"],
		},
	}

	responses.JSON(w, http.StatusOK, res)
}
