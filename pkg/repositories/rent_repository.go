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
	"go.mongodb.org/mongo-driver/mongo/options"
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
			"clientName":    rent.ClientName,
			"bookAuthor":    rent.BookAuthor,
			"bookPublisher": rent.BookPublisher,
			"bookName":      rent.BookName,
			"delivered":     bson.M{"$eq": false},
		})

	if err != nil {
		return nil, err
	}
	//Verificando se já existe algum aluguel aberto para o livro do mesmo usuário
	if existRent > 0 {
		return nil, errors.New("there is already an open rent for this book in this client's name")
	}

	//Verificando se o usuário existe
	clientRepository := CreateNewClientRepository()
	if _, err := clientRepository.GetClientByName(rent.ClientName); err != nil {
		return nil, errors.New("client not found")
	}
	//Verificando se o livro existe
	bookRepository := CreateNewBooksRepository()
	book, err := bookRepository.GetBookByName(rent.BookName)
	if err != nil {
		return nil, errors.New("book not found")
	}
	//Verificando se a editora existe
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
func (repository *Rents) GetNotDeliveredRents(options *options.FindOptions) ([]entities.Rent, error) {
	cur, err := repository.collection.Find(repository.ctx, bson.M{"delivered": false}, options)
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
	fil := bson.M{
		"clientName":    bson.M{"$regex": filters["clientName"], "$options": "i"},
		"bookAuthor":    bson.M{"$regex": filters["bookAuthor"], "$options": "i"},
		"bookPublisher": bson.M{"$regex": filters["bookPublisher"], "$options": "i"},
		"bookName":      bson.M{"$regex": filters["bookName"], "$options": "i"},
	}
	if filters["rentStatus"] != "" {
		delivered, _ := strconv.ParseBool(filters["rentStatus"])
		fil["delivered"] = bson.M{"$eq": delivered}
	}

	cur, err := repository.collection.Find(repository.ctx, fil)
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

	if len(rents) == 0 {
		return nil, errors.New("there are no rents")
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

func (repository *Rents) CheckRentedBooksByClientName(clientName string) (bool, error) {
	exists, err := repository.collection.CountDocuments(repository.ctx, bson.M{"clientName": clientName, "delivered": false})
	if err != nil {
		return false, err
	}

	return exists > 0, nil
}

func (repository *Rents) GetNumberRentsByClient(clientName string) (map[string]interface{}, error) {
	numRents, err := repository.collection.CountDocuments(repository.ctx, bson.M{"clientName": clientName})
	if err != nil {
		return nil, err
	}

	if numRents == 0 {
		return nil, errors.New("this client has no rentals")
	}

	res := map[string]interface{}{
		"clientName":  clientName,
		"NumberRents": numRents,
	}

	return res, nil
}

func (repository *Rents) GetReturnedBooks() (map[string]interface{}, error) {
	var outTheTimeLimit, withinTheTimeLimit []entities.Rent
	cur, err := repository.collection.Find(repository.ctx, bson.M{"delivered": bson.M{"$eq": true}})
	if err != nil {
		return nil, err
	}

	defer cur.Close(repository.ctx)

	for cur.Next(repository.ctx) {
		var rent entities.Rent
		err := cur.Decode(&rent)
		if err != nil {
			return nil, err
		}
		if rent.ExpectedDeliveryDate.Time.Equal(rent.DeliveryDate.Time) {
			withinTheTimeLimit = append(withinTheTimeLimit, rent)
		} else if rent.ExpectedDeliveryDate.Time.Before(rent.DeliveryDate.Time) {
			outTheTimeLimit = append(outTheTimeLimit, rent)
		} else {
			withinTheTimeLimit = append(withinTheTimeLimit, rent)
		}
	}
	res := map[string]interface{}{
		"outTheTimeLimit":    utils.ReturnMessageOrValue(outTheTimeLimit, "out"),
		"withinTheTimeLimit": utils.ReturnMessageOrValue(withinTheTimeLimit, "within"),
	}
	return res, nil
}

func (repository *Rents) GetRentedBooks() (map[string]interface{}, error) {
	numRents, err := repository.collection.CountDocuments(repository.ctx, bson.M{"delivered": false})
	if err != nil {
		return nil, err
	}

	if numRents == 0 {
		return nil, errors.New("there are not rented books")
	}

	res := map[string]interface{}{
		"NumberOfRentedBooks": numRents,
	}

	return res, nil
}

func (repository *Rents) GetNumberOfOverdueBooks() (map[string]interface{}, error) {
	timeNow := utils.CustomTime{Time: time.Now()}
	timeNow.Format()
	cur, err := repository.collection.Find(repository.ctx, bson.M{
		"$and": bson.A{
			bson.M{"expectedDeliveryDate.time": bson.M{"$lt": timeNow.Time}},
			bson.M{"delivered": false},
		},
	})
	if err != nil {
		return nil, err
	}

	defer cur.Close(repository.ctx)

	var outTheTimeLimit []entities.Rent
	if err := cur.All(repository.ctx, &outTheTimeLimit); err != nil {
		return nil, err
	}

	if len(outTheTimeLimit) == 0 {
		return nil, errors.New("there are no overdue books")
	}

	res := map[string]interface{}{
		"NumberOfOverdueBooks": len(outTheTimeLimit),
	}

	return res, nil
}

func (repository *Rents) GetNumberOfBooksRentsByClient(clientName string) (map[string]interface{}, error) {
	pipe := []bson.M{
		{
			"$match": bson.M{
				"clientName": clientName,
			},
		},
		{
			"$group": bson.M{
				"_id":      "$bookName",
				"numBooks": bson.M{"$sum": 1},
			},
		},
	}
	var mostRentedBooks []map[string]interface{}
	cur, err := repository.collection.Aggregate(repository.ctx, pipe)
	if err != nil {
		return nil, err
	}
	defer cur.Close(repository.ctx)

	if err := cur.All(repository.ctx, &mostRentedBooks); err != nil {
		return nil, err
	}

	if len(mostRentedBooks) == 0 {
		return nil, errors.New("this client has no rentals")
	}

	res := map[string]interface{}{
		"clientName": clientName,
		"books":      mostRentedBooks,
	}

	return res, nil
}
