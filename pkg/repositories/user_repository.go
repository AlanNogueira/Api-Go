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

type Users struct {
	collection *mongo.Collection
	ctx        context.Context
}

func CreateNewUserRepository() *Users {
	collection := configuration.Client.Database(configuration.DBName).Collection("users")
	return &Users{collection, context.Background()}
}

func (repository *Users) Create(user entities.User) (map[string]interface{}, error) {
	exists, _ := repository.collection.CountDocuments(repository.ctx, bson.M{"name": user.Name})
	if exists > 0 {
		return nil, errors.New("this user already exists")
	}

	req, err := repository.collection.InsertOne(repository.ctx, user)
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

func (repository *Users) RemoveUser(userId string) (map[string]interface{}, error) {
	var user entities.User
	idPrimitive, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New("invalid Id")
	}

	err = repository.collection.FindOne(repository.ctx, bson.M{"_id": idPrimitive}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}

	rentRepository := CreateNewRentRepository()
	existsRentedBooks, err := rentRepository.CheckRentedBooksByUserName(user.Name)
	if err != nil {
		return nil, errors.New("unable to perform this action at the moment, please try again later")
	}
	if existsRentedBooks {
		return nil, errors.New("this user has rented books, you can't delete it")
	}

	_, err = repository.collection.DeleteOne(repository.ctx, bson.M{"_id": idPrimitive})
	if err != nil {
		return nil, errors.New("unable to perform this action at the moment, please try again later")
	}

	res := map[string]interface{}{
		"data": "User deleted socessfully",
	}

	return res, nil
}

func (repository *Users) GetUser(userName string) ([]entities.User, error) {
	cur, err := repository.collection.Find(repository.ctx, bson.M{"name": bson.M{"$regex": userName, "$options": "i"}})
	if err != nil {
		return nil, err
	}

	defer cur.Close(repository.ctx)

	var users []entities.User

	for cur.Next(repository.ctx) {
		var user entities.User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, errors.New("there are no registered books with this name")
	}

	return users, nil
}

func (repository *Users) GetUsers() ([]entities.User, error) {
	cur, err := repository.collection.Find(repository.ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(repository.ctx)

	var users []entities.User

	for cur.Next(repository.ctx) {
		var user entities.User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, errors.New("there are no registered users")
	}

	return users, nil
}

func (repository *Users) UpdateUser(userId string, userNewData entities.User) (map[string]interface{}, error) {
	var user entities.User
	idPrimitive, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New("invalid Id")
	}

	err = repository.collection.FindOne(repository.ctx, bson.M{"_id": idPrimitive}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}

	filter := bson.M{"_id": idPrimitive}
	fields := bson.M{"$set": userNewData}

	_, err = repository.collection.UpdateOne(repository.ctx, filter, fields)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"data": "User updated successfully.",
	}

	return res, nil
}

func (repository *Users) GetUserByName(userName string) (entities.User, error) {
	var user entities.User

	err := repository.collection.FindOne(repository.ctx, bson.M{"name": userName}).Decode(&user)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
