package entities

import (
	"Api-Go/pkg/entities/utils"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type RentalCompany struct {
	Id       string `bson:"_id,omitempty"`
	Name     string `bson:"name,omitempty"`
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}

func (rentalCompany *RentalCompany) Prepare() error {
	if err := rentalCompany.validate(); err != nil {
		return err
	}

	rentalCompany.format()
	return nil
}
func (rentalCompany *RentalCompany) validate() error {
	rentalCompanyErrors := ""

	if rentalCompany.Name == "" || !utils.ValidateName(rentalCompany.Name) {
		rentalCompanyErrors += "name, "
	}
	if rentalCompany.Email == "" || !utils.ValidateEmail(rentalCompany.Email) {
		rentalCompanyErrors += "email, "
	}
	if rentalCompany.Password == "" || !utils.ValidatePassword(rentalCompany.Password) {
		rentalCompanyErrors += "password - rules: minimun 8 characters, at least 1 number, at least 1 upper case, at least 1 special character , "
	}
	if rentalCompanyErrors != "" {
		return errors.New("There are errors in the fields: " + rentalCompanyErrors)
	}
	return nil
}

func (rentalCompany *RentalCompany) format() {
	rentalCompany.Name = strings.TrimSpace(rentalCompany.Name)
	rentalCompany.Email = strings.TrimSpace(rentalCompany.Email)
}

func (rentalCompany *RentalCompany) GetHash() (err error) {
	password := []byte(rentalCompany.Password)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	rentalCompany.Password = string(hash)
	if err != nil {
		return
	}

	return nil
}
