package configuration

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	uri    = "mongodb://localhost:27017"
	dbName = "shop_db"
)

func DbConnect(collectionName string) (*mongo.Collection, error) {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)

	return collection, nil
}
