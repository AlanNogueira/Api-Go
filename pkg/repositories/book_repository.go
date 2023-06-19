package repositories

import (
	"Api-Go/pkg/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Books struct {
	collection *mongo.Collection
	ctx        context.Context
}

func CreateNewBooksRepository(collection *mongo.Collection) *Books {
	return &Books{collection, context.Background()}
}

func (repository *Books) Create(book models.Book) (map[string]interface{}, error) {
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
	var book models.Book
	idPrimitive, err := primitive.ObjectIDFromHex(bookId)
	if err != nil {
		return nil, err
	}

	err = repository.collection.FindOne(repository.ctx, bson.M{"_id": idPrimitive}).Decode(&book)
	if err != nil {
		return nil, errors.New("book not found")
	}

	_, err = repository.collection.DeleteOne(repository.ctx, bson.M{"_id": idPrimitive})
	if err != nil {
		return nil, errors.New("error deleting book")
	}

	res := map[string]interface{}{
		"data": "Document deleted socessfully",
	}

	return res, nil
}

func (repository *Books) GetBook(bookId string) (models.Book, error) {
	var book models.Book
	idPrimitive, err := primitive.ObjectIDFromHex(bookId)
	if err != nil {
		return models.Book{}, err
	}

	err = repository.collection.FindOne(repository.ctx, bson.M{"_id": idPrimitive}).Decode(&book)
	if err != nil {
		return models.Book{}, errors.New("book not found")
	}

	return book, nil
}

func (repository *Books) GetBooks() ([]models.Book, error) {
	cur, err := repository.collection.Find(repository.ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(repository.ctx)
	var books []models.Book

	for cur.Next(repository.ctx) {
		var book models.Book

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

func (repository *Books) UpdateBook(bookId string, bookNewData models.Book) (map[string]interface{}, error) {
	var book models.Book
	idPrimitive, err := primitive.ObjectIDFromHex(bookId)
	if err != nil {
		return nil, errors.New("invalid Id")
	}

	err = repository.collection.FindOne(repository.ctx, bson.M{"_id": idPrimitive}).Decode(&book)
	if err != nil {
		return nil, errors.New("book not found")
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
