package repositories

import (
	"Api-Go/pkg/configuration"
	"Api-Go/pkg/entities"
	"Api-Go/pkg/entities/utils"
	"context"
	"errors"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Rents struct {
	collection *mongo.Collection
	ctx        context.Context
}

func CreateNewRentRepository() *Rents {
	collection := configuration.Client.Database(configuration.DBName).Collection("rents")
	return &Rents{collection, context.Background()}
}

func (repository *Rents) CreateRent(rent entities.Rent) (map[string]interface{}, error) {
	existRent, err := repository.collection.CountDocuments(repository.ctx,
		bson.M{
			"userName":      bson.M{"$regex": rent.UserName, "$options": "i"},
			"bookAuthor":    bson.M{"$regex": rent.BookAuthor, "$options": "i"},
			"bookPublisher": bson.M{"$regex": rent.BookPublisher, "$options": "i"},
			"bookName":      bson.M{"$regex": rent.BookName, "$options": "i"},
			"delivered":     bson.M{"$eq": false},
		})

	if err != nil {
		return nil, err
	}
	//Verificando se já existe algum aluguel aberto para o livro do mesmo usuário
	if existRent > 0 {
		return nil, errors.New("there is already an open rent for this book in this user's name")
	}

	//Verificando se o usuário existe
	userRepository := CreateNewUserRepository()
	if _, err := userRepository.GetUserByName(rent.UserName); err != nil {
		return nil, errors.New("user not found")
	}
	//Verificando se o livro existe
	bookRepository := CreateNewBooksRepository()
	book, err := bookRepository.GetBookByName(rent.BookName)
	if err != nil {
		return nil, errors.New("book not found")
	}
	//Verificando se o autor existe
	publisherRepository := CreateNewPublishersRepository()
	if _, err := publisherRepository.GetPublisherByName(rent.BookPublisher); err != nil {
		return nil, errors.New("publisher not found")
	}

	if err := bookRepository.AddRentOfBook(book); err != nil {
		return nil, err
	}

	req, err := repository.collection.InsertOne(repository.ctx, rent)
	if err != nil {
		return nil, err
	}

	insertedId := req.InsertedID

	res := map[string]interface{}{
		"data": map[string]interface{}{
			"insertedId": insertedId,
		},
	}

	return res, nil
}
func (repository *Rents) GetNotDeliveredRents() ([]entities.Rent, error) {
	cur, err := repository.collection.Find(repository.ctx, bson.M{"delivered": false})
	if err != nil {
		return nil, err
	}

	defer cur.Close(repository.ctx)

	var rents []entities.Rent

	for cur.Next(repository.ctx) {
		var rent entities.Rent
		if cur.Decode(&rent) != nil {
			return nil, err
		}

		rents = append(rents, rent)
	}

	if len(rents) == 0 {
		return nil, errors.New("there are no unfinalized rent")
	}

	return rents, nil
}

func (repository *Rents) GetRent(filters map[string]string) ([]entities.Rent, error) {
	delivered, _ := strconv.ParseBool(filters["rentStatus"])
	cur, err := repository.collection.Find(repository.ctx,
		bson.M{
			"userName":      bson.M{"$regex": filters["userName"], "$options": "i"},
			"bookAuthor":    bson.M{"$regex": filters["bookAuthor"], "$options": "i"},
			"bookPublisher": bson.M{"$regex": filters["bookPublisher"], "$options": "i"},
			"bookName":      bson.M{"$regex": filters["bookName"], "$options": "i"},
			"delivered":     bson.M{"$eq": delivered},
		})
	if err != nil {
		return nil, err
	}

	defer cur.Close(repository.ctx)

	var rents []entities.Rent

	for cur.Next(repository.ctx) {
		var rent entities.Rent
		err := cur.Decode(&rent)
		if err != nil {
			return nil, err
		}

		rents = append(rents, rent)
	}

	return rents, nil
}

func (repository *Rents) FinalizeRent(rentId string) (map[string]interface{}, error) {
	idPrimitive, err := primitive.ObjectIDFromHex(rentId)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	var rent entities.Rent
	err = repository.collection.FindOne(repository.ctx, bson.M{"_id": idPrimitive}).Decode(&rent)
	if err != nil {
		return nil, errors.New("rent not found")
	}

	if rent.Delivered {
		return nil, errors.New("rent already finalized")
	}

	bookRepository := CreateNewBooksRepository()
	book, err := bookRepository.GetBookByName(rent.BookName)
	if err != nil {
		return nil, errors.New("book not found")
	}

	if err := bookRepository.FinalizeBookRent(book); err != nil {
		return nil, err
	}

	newDataRent := entities.Rent{
		Delivered:    true,
		DeliveryDate: utils.CustomTime{Time: time.Now()},
	}
	newDataRent.DeliveryDate.Format()

	filter := bson.M{"_id": idPrimitive}
	fields := bson.M{"$set": newDataRent}

	_, err = repository.collection.UpdateOne(repository.ctx, filter, fields)

	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"Data": "Rent finalized successfully",
	}

	return res, nil
}

func (repository *Rents) CheckRentedBooksByUserName(userName string) (bool, error) {
	exists, err := repository.collection.CountDocuments(repository.ctx, bson.M{"userName": userName, "delivered": false})
	if err != nil {
		return false, err
	}

	return exists > 0, nil
}
