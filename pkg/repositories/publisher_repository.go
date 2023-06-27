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

type Publishers struct {
	collection *mongo.Collection
	ctx        context.Context
}

func CreateNewPublishersRepository() *Publishers {
	collection := configuration.Client.Database(configuration.DBName).Collection("publishers")
	return &Publishers{collection, context.Background()}
}

func (repository *Publishers) CreatePublisher(publisher entities.Publisher) (map[string]interface{}, error) {
	exists, _ := repository.collection.CountDocuments(repository.ctx, bson.M{"name": publisher.Name})
	if exists > 0 {
		return nil, errors.New("this publisher already exists")
	}

	req, err := repository.collection.InsertOne(repository.ctx, publisher)
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
func (repository *Publishers) RemovePublisher(publisherId string) (map[string]interface{}, error) {
	var publisher entities.Publisher
	idPrimitive, err := primitive.ObjectIDFromHex(publisherId)
	if err != nil {
		return nil, errors.New("invalid Id")
	}

	err = repository.collection.FindOne(repository.ctx, bson.M{"_id": idPrimitive}).Decode(&publisher)
	if err != nil {
		return nil, errors.New("publisher not found")
	}

	_, err = repository.collection.DeleteOne(repository.ctx, bson.M{"_id": idPrimitive})
	if err != nil {
		return nil, errors.New("unable to perform this action at the moment, please try again later")
	}

	res := map[string]interface{}{
		"data": "Publisher deleted successfully",
	}

	return res, nil
}
func (repository *Publishers) GetPublisher(publisherName string) ([]entities.Publisher, error) {
	cur, err := repository.collection.Find(repository.ctx, bson.M{"name": bson.M{"$regex": publisherName, "$options": "i"}})
	if err != nil {
		return nil, err
	}

	defer cur.Close(repository.ctx)

	var publishers []entities.Publisher

	for cur.Next(repository.ctx) {
		var publisher entities.Publisher

		if err := cur.Decode(&publisher); err != nil {
			return nil, err
		}

		publishers = append(publishers, publisher)
	}

	if len(publishers) == 0 {
		return nil, errors.New("there are no registered publisher with this name")
	}

	return publishers, nil
}

func (repository *Publishers) GetPublishers(options *options.FindOptions) ([]entities.Publisher, error) {
	cur, err := repository.collection.Find(repository.ctx, bson.D{}, options)
	if err != nil {
		return nil, err
	}

	defer cur.Close(repository.ctx)

	var publishers []entities.Publisher

	for cur.Next(repository.ctx) {
		var publisher entities.Publisher

		if err := cur.Decode(&publisher); err != nil {
			return nil, err
		}

		publishers = append(publishers, publisher)
	}

	if len(publishers) == 0 {
		return nil, errors.New("there are no registered publishers")
	}

	return publishers, nil
}

func (repository *Publishers) UpdatePublisher(publisherId string, newPublisherData entities.Publisher) (map[string]interface{}, error) {
	var publisher entities.Publisher
	idPrimitive, err := primitive.ObjectIDFromHex(publisherId)
	if err != nil {
		return nil, errors.New("invalid Id")
	}

	err = repository.collection.FindOne(repository.ctx, bson.M{"_id": idPrimitive}).Decode(&publisher)
	if err != nil {
		return nil, errors.New("publisher not found")
	}

	filter := bson.M{"_id": idPrimitive}
	fields := bson.M{"$set": newPublisherData}

	_, err = repository.collection.UpdateOne(repository.ctx, filter, fields)
	if err != nil {
		return nil, errors.New("unable to perform this action at the moment, please try again later")
	}

	res := map[string]interface{}{
		"data": "Publisher updated successfully",
	}

	return res, nil
}

func (repository *Publishers) GetPublisherByName(publisherName string) (entities.Publisher, error) {
	var publisher entities.Publisher

	err := repository.collection.FindOne(repository.ctx, bson.M{"name": publisherName}).Decode(&publisher)
	if err != nil {
		return entities.Publisher{}, err
	}

	return publisher, nil
}
