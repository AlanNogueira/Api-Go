package repositories

import (
	"Api-Go/pkg/configuration"
	"Api-Go/pkg/entities"
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

func CreateNewBooksRepository() (*Books, error) {
	collection, err := configuration.DbConnect("books")
	if err != nil {
		return nil, err
	}

	return &Books{collection, context.Background()}, nil
}

func (repository *Books) Create(book entities.Book) (map[string]interface{}, error) {
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
		bson.M{"name": bson.M{"$regex": bookData.Name, "$options": "i"},
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

func (repository *Books) GetBooks() ([]entities.Book, error) {
	cur, err := repository.collection.Find(repository.ctx, bson.D{})
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
