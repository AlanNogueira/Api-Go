package repositories

import (
	"Api-Go/pkg/auth"
	"Api-Go/pkg/configuration"
	"Api-Go/pkg/entities"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type RentalCompanyRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func CreateNewRentalCompanyRepository() *RentalCompanyRepository {
	collection := configuration.Client.Database(configuration.DBName).Collection("rental_companies")
	return &RentalCompanyRepository{collection, context.Background()}
}

func (repository *RentalCompanyRepository) CreateRentalCompany(rentalCompany entities.RentalCompany) (map[string]interface{}, error) {
	exists, _ := repository.collection.CountDocuments(repository.ctx, bson.M{"email": rentalCompany.Email, "name": rentalCompany.Name})
	if exists > 0 {
		return nil, errors.New("this rental company already exists")
	}

	if err := rentalCompany.GetHash(); err != nil {
		return nil, err
	}

	req, err := repository.collection.InsertOne(repository.ctx, rentalCompany)
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

func (repository *RentalCompanyRepository) AuthenticateUser(rentalCompany entities.RentalCompany) (map[string]interface{}, error) {
	var dbRentalCompany entities.RentalCompany
	err := repository.collection.FindOne(repository.ctx, bson.M{"email": rentalCompany.Email}).Decode(&dbRentalCompany)
	if err != nil {
		return nil, errors.New("rental company not found")
	}

	dbPassword := []byte(dbRentalCompany.Password)
	password := []byte(rentalCompany.Password)
	if err := bcrypt.CompareHashAndPassword(dbPassword, password); err != nil {
		return nil, errors.New("invalid password")
	}

	token, err := auth.CreateToken(dbRentalCompany.Id)
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
