package repositories

import (
	"Api-Go/pkg/auth"
	"Api-Go/pkg/configuration"
	"Api-Go/pkg/entities"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func CreateNewUserRepository() *UserRepository {
	collection := configuration.Client.Database(configuration.DBName).Collection("users")
	return &UserRepository{collection, context.Background()}
}

func (repository *UserRepository) CreateUser(user entities.User) (map[string]interface{}, error) {
	exists, _ := repository.collection.CountDocuments(repository.ctx, bson.M{"email": user.Email})
	if exists > 0 {
		return nil, errors.New("a user with this email already exists")
	}

	if err := user.GetHash(); err != nil {
		return nil, err
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

func (repository *UserRepository) GetUsers(options *options.FindOptions) ([]entities.User, error) {
	var users []entities.User
	options.SetProjection(bson.M{"password": 0})
	cur, err := repository.collection.Find(repository.ctx, bson.M{}, options)
	if err != nil {
		return nil, err
	}

	if err := cur.All(repository.ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (repository *UserRepository) AuthenticateUser(user entities.User) (map[string]interface{}, error) {
	var dbUser entities.User
	err := repository.collection.FindOne(repository.ctx, bson.M{"email": user.Email}).Decode(&dbUser)
	if err != nil {
		return nil, errors.New("user not found")
	}

	dbPassword := []byte(dbUser.Password)
	password := []byte(user.Password)
	if err := bcrypt.CompareHashAndPassword(dbPassword, password); err != nil {
		return nil, errors.New("invalid password")
	}

	token, err := auth.CreateToken(dbUser.Id)
	if err != nil {
		return nil, errors.New("error creating token")
	}

	res := map[string]interface{}{
		"data": map[string]interface{}{
			"Message": "User Authenticated",
			"Token: ": token,
		},
	}

	return res, nil
}
