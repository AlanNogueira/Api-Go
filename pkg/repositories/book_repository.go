package repositories

import (
	"Api-Go/pkg/configuration"
	"Api-Go/pkg/entities"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Books struct {
	collection *mongo.Collection
	ctx        context.Context
}

func CreateNewBooksRepository() *Books {
	collection := configuration.Client.Database(configuration.DBName).Collection("books")
	return &Books{collection, context.Background()}
}

func (repository *Books) Create(book entities.Book) (map[string]interface{}, error) {
	exists, _ := repository.collection.CountDocuments(repository.ctx, bson.M{"name": book.Name})
	if exists > 0 {
		return nil, errors.New("this book already exists")
	}

	if book.Stock == 0 {
		return nil, errors.New("stock cannot be 0")
	}

	req, err := repository.collection.InsertOne(repository.ctx, book)
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

func (repository *Books) RemoveBook(bookId string) (map[string]interface{}, error) {
	var book entities.Book
	idPrimitive, err := primitive.ObjectIDFromHex(bookId)
	if err != nil {
		return nil, err
	}

	err = repository.collection.FindOne(repository.ctx, bson.M{"_id": idPrimitive}).Decode(&book)
	if err != nil {
		return nil, errors.New("book not found")
	}

	if book.Rented != 0 {
		return nil, errors.New("unable to delete rented books")
	}

	_, err = repository.collection.DeleteOne(repository.ctx, bson.M{"_id": idPrimitive})
	if err != nil {
		return nil, errors.New("error deleting book")
	}

	res := map[string]interface{}{
		"data": "Book deleted socessfully",
	}

	return res, nil
}

func (repository *Books) GetBook(bookData entities.Book) ([]entities.Book, error) {
	cur, err := repository.collection.Find(repository.ctx,
		bson.M{
			"name":             bson.M{"$regex": bookData.Name, "$options": "i"},
			"author":           bson.M{"$regex": bookData.Author, "$options": "i"},
			"publisher":        bson.M{"$regex": bookData.Publisher, "$options": "i"},
			"releaseDate.time": bson.M{"$eq": bookData.ReleaseDate.Time}})
	if err != nil {
		return nil, err
	}

	defer cur.Close(repository.ctx)

	var books []entities.Book

	for cur.Next(repository.ctx) {
		var book entities.Book

		if err := cur.Decode(&book); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if len(books) == 0 {
		return nil, errors.New("there are no registered books")
	}

	return books, nil
}

func (repository *Books) GetBooks(options *options.FindOptions) ([]entities.Book, error) {
	cur, err := repository.collection.Find(repository.ctx, bson.D{}, options)
	if err != nil {
		return nil, err
	}

	defer cur.Close(repository.ctx)

	var books []entities.Book

	for cur.Next(repository.ctx) {
		var book entities.Book

		if err := cur.Decode(&book); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if len(books) == 0 {
		return nil, errors.New("there are no registered books")
	}

	return books, nil
}

func (repository *Books) UpdateBook(bookId string, bookNewData entities.Book) (map[string]interface{}, error) {
	var book entities.Book
	idPrimitive, err := primitive.ObjectIDFromHex(bookId)
	if err != nil {
		return nil, errors.New("invalid Id")
	}

	if bookNewData.Rented != 0 {
		return nil, errors.New("unable to update the number of rented books")
	}

	err = repository.collection.FindOne(repository.ctx, bson.M{"_id": idPrimitive}).Decode(&book)
	if err != nil {
		return nil, errors.New("book not found")
	}

	if bookNewData.Stock > 0 && bookNewData.Stock < book.Rented {
		return nil, errors.New("it is not possible to update the stock for a quantity smaller than the rented")
	}

	filter := bson.M{"_id": idPrimitive}
	fields := bson.M{"$set": bookNewData}

	_, err = repository.collection.UpdateOne(repository.ctx, filter, fields)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"data": "Book updated successfully.",
	}

	return res, nil
}

func (repository *Books) GetBookByName(bookName string) (entities.Book, error) {
	var book entities.Book

	err := repository.collection.FindOne(repository.ctx, bson.M{"name": bookName}).Decode(&book)
	if err != nil {
		return entities.Book{}, errors.New("book not found")
	}

	return book, nil
}

func (repository *Books) AddRentOfBook(book entities.Book) error {
	newRentedNumber := book.Rented + 1
	numberOfRents := book.NumberOfRents + 1

	if newRentedNumber > book.Stock {
		return errors.New("there is no more of this book in stock")
	}

	filter := bson.M{"name": book.Name}
	fields := bson.M{"$set": bson.M{"rented": newRentedNumber, "numberOfRents": numberOfRents}}

	_, err := repository.collection.UpdateOne(repository.ctx, filter, fields)
	if err != nil {
		return err
	}

	return nil
}

func (repository *Books) FinalizeBookRent(book entities.Book) error {
	if book.Rented == 0 {
		return errors.New("there is no rent to finalize")
	}

	newRentedNumber := book.Rented - 1
	filter := bson.M{"name": book.Name}
	fields := bson.M{"$set": bson.M{"rented": newRentedNumber}}

	_, err := repository.collection.UpdateOne(repository.ctx, filter, fields)
	if err != nil {
		return err
	}

	return nil
}

func (repository *Books) GetMostRentedBook() (map[string]interface{}, error) {
	var mostRentedBook []map[string]interface{}

	pipe := []bson.M{
		{
			"$sort": bson.M{"numberOfRents": -1},
		},
	}
	cur, err := repository.collection.Aggregate(repository.ctx, pipe)
	if err != nil {
		return nil, err
	}

	defer cur.Close(repository.ctx)
	if err := cur.All(repository.ctx, &mostRentedBook); err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"mostRentedBook": mostRentedBook[0],
	}

	return res, nil
}
