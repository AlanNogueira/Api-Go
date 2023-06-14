package repositories

import (
	"Api-Go/pkg/models"
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

func CreateNewUserRepository(collection *mongo.Collection) *Users {
	return &Users{collection, context.Background()}
}

func (repository *Users) Create(user models.User) (map[string]interface{}, error) {
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
	var user models.User
	idPrimitive, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New("invalid Id")
	}

	err = repository.collection.FindOne(repository.ctx, bson.M{"_id": idPrimitive}).Decode(&user)
	if err != nil {
		return nil, errors.New("User not found")
	}

	_, err = repository.collection.DeleteOne(repository.ctx, bson.M{"_id": idPrimitive})
	if err != nil {
		return nil, errors.New("unable to perform this action at the moment, please try again later")
	}

	res := map[string]interface{}{
		"data": "Document deleted socessfully",
	}

	return res, nil
}

func (repository *Users) GetUser(userId string) (map[string]interface{}, error) {
	user := map[string]interface{}{}
	idPrimitive, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New("invalid Id")
	}

	err = repository.collection.FindOne(repository.ctx, bson.M{"_id": idPrimitive}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (repository *Users) GetUsers() (map[string]interface{}, error) {
	cur, err := repository.collection.Find(repository.ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(repository.ctx)

	var products []bson.M

	for cur.Next(repository.ctx) {
		var product bson.M
		if err := cur.Decode(&product); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if len(products) == 0 {
		return nil, errors.New("there are no registered users")
	}

	users := map[string]interface{}{
		"data": products,
	}

	return users, nil
}

func (repository *Users) UpdateUser(userId string, userNewData models.User) (map[string]interface{}, error) {
	user := map[string]interface{}{}
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
