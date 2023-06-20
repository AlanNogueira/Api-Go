package configuration

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

const (
	uri    = "mongodb://localhost:27017"
	DBName = "library_db"
)

func DbConnect() *mongo.Client {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func init() {
	Client = DbConnect()
}
