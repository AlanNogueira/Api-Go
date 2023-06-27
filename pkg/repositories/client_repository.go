package repositories

import (
	"Api-Go/pkg/configuration"
	"Api-Go/pkg/entities"
	"context"
	"errors"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Clients struct {
	collection *mongo.Collection
	ctx        context.Context
}

func CreateNewClientRepository() *Clients {
	collection := configuration.Client.Database(configuration.DBName).Collection("clients")
	return &Clients{collection, context.Background()}
}

func (repository *Clients) Create(client entities.Client) (map[string]interface{}, error) {
	exists, _ := repository.collection.CountDocuments(repository.ctx, bson.M{"taxNumber": client.TaxNumber})
	if exists > 0 {
		return nil, errors.New("there is already a registered client with this cpf")
	}

	exists, _ = repository.collection.CountDocuments(repository.ctx, bson.M{"name": client.Name})
	if exists > 0 {
		return nil, errors.New("this clientName is already in use")
	}

	req, err := repository.collection.InsertOne(repository.ctx, client)
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

func (repository *Clients) RemoveClient(clientId string) (map[string]interface{}, error) {
	var client entities.Client
	idPrimitive, err := primitive.ObjectIDFromHex(clientId)
	if err != nil {
		return nil, errors.New("invalid Id")
	}

	err = repository.collection.FindOne(repository.ctx, bson.M{"_id": idPrimitive}).Decode(&client)
	if err != nil {
		return nil, errors.New("client not found")
	}

	rentRepository := CreateNewRentRepository()
	existsRentedBooks, err := rentRepository.CheckRentedBooksByClientName(client.Name)
	if err != nil {
		return nil, errors.New("unable to perform this action at the moment, please try again later")
	}
	if existsRentedBooks {
		return nil, errors.New("this client has rented books, you can't delete it")
	}

	_, err = repository.collection.DeleteOne(repository.ctx, bson.M{"_id": idPrimitive})
	if err != nil {
		return nil, errors.New("unable to perform this action at the moment, please try again later")
	}

	res := map[string]interface{}{
		"data": "client deleted socessfully",
	}

	return res, nil
}

func (repository *Clients) GetClient(taxNumber string) (entities.Client, error) {
	taxNumber = regexp.MustCompile(`[^0-9]+`).ReplaceAllString(taxNumber, "")
	var client entities.Client
	err := repository.collection.FindOne(repository.ctx, bson.M{"taxNumber": taxNumber}).Decode(&client)
	if err != nil {
		return entities.Client{}, errors.New("client not found")
	}

	return client, nil
}

func (repository *Clients) GetClients(options *options.FindOptions) ([]entities.Client, error) {
	cur, err := repository.collection.Find(repository.ctx, bson.D{}, options)
	if err != nil {
		return nil, err
	}

	defer cur.Close(repository.ctx)

	var clients []entities.Client

	for cur.Next(repository.ctx) {
		var client entities.Client
		if err := cur.Decode(&client); err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	if len(clients) == 0 {
		return nil, errors.New("there are no registered clients")
	}

	return clients, nil
}

func (repository *Clients) UpdateClient(clientId string, clientNewData entities.Client) (map[string]interface{}, error) {
	var client entities.Client
	idPrimitive, err := primitive.ObjectIDFromHex(clientId)
	if err != nil {
		return nil, errors.New("invalid Id")
	}

	err = repository.collection.FindOne(repository.ctx, bson.M{"_id": idPrimitive}).Decode(&client)
	if err != nil {
		return nil, errors.New("client not found")
	}

	filter := bson.M{"_id": idPrimitive}
	fields := bson.M{"$set": clientNewData}

	_, err = repository.collection.UpdateOne(repository.ctx, filter, fields)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"data": "client updated successfully.",
	}

	return res, nil
}

func (repository *Clients) GetClientByName(clientName string) (entities.Client, error) {
	var client entities.Client

	err := repository.collection.FindOne(repository.ctx, bson.M{"name": clientName}).Decode(&client)
	if err != nil {
		return entities.Client{}, errors.New("client not found")
	}

	return client, nil
}
