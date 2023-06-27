package entities

import (
	"Api-Go/pkg/entities/utils"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `bson:"_id,omitempty"`
	Name     string `bson:"name,omitempty"`
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}

func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.format()
	return nil
}
func (user *User) validate() error {
	userErrors := ""

	if user.Name == "" || !utils.ValidateName(user.Name) {
		userErrors += "name, "
	}
	if user.Email == "" || !utils.ValidateEmail(user.Email) {
		userErrors += "email, "
	}
	if user.Password == "" || !utils.ValidatePassword(user.Password) {
		userErrors += "password - rules: minimun 8 characters, at least 1 number, at least 1 upper case, at least 1 special character , "
	}
	if userErrors != "" {
		return errors.New("There are errors in the fields: " + userErrors)
	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
}

func (user *User) GetHash() (err error) {
	password := []byte(user.Password)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	user.Password = string(hash)
	if err != nil {
		return
	}

	return nil
}
