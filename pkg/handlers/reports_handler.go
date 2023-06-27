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

func GetNumberOfBooksRentsByClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	clientName := params["clientName"]
	rentsRepository := repositories.CreateNewRentRepository()
	response, err := rentsRepository.GetNumberOfBooksRentsByClient(clientName)
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

func GetNumberRentsByClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	clientName := params["clientName"]

	rentRepository := repositories.CreateNewRentRepository()
	reponse, err := rentRepository.GetNumberRentsByClient(clientName)
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

	var rentedBooksByClient []map[string]interface{}
	var numRentsByClient []map[string]interface{}
	clientRepository := repositories.CreateNewClientRepository()
	options := options.Find()
	clients, err := clientRepository.GetClients(options)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	for i := 0; i < len(clients); i++ {
		response, err := rentsRepository.GetNumberOfBooksRentsByClient(clients[i].Name)
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}
		rentedBooksByClient = append(rentedBooksByClient, response)

		reponse, err := rentsRepository.GetNumberRentsByClient(clients[i].Name)
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}
		numRentsByClient = append(numRentsByClient, reponse)
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
			"NumberOfRentedBooks":        rentedBooks["NumberOfRentedBooks"],
			"NumberOfOverdueBooks":       overdueBooks["NumberOfOverdueBooks"],
			"NumberOfBooksRentsByClient": rentedBooksByClient,
			"ReturnedBooks":              returnedBooks,
			"NumberRentsByClient":        numRentsByClient,
			"MostRentedBook":             mostRentedBook["mostRentedBook"],
		},
	}

	responses.JSON(w, http.StatusOK, res)
}
