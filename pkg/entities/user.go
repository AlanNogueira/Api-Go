package entities

import (
	"Api-Go/pkg/entities/utils"
	"errors"
	"regexp"
	"strings"
)

type User struct {
	Id        string        `bson:"_id,omitempty"`
	Name      string        `bson:"name,omitempty"`
	Email     string        `bson:"email,omitempty"`
	Phone     string        `bson:"phone,omitempty"`
	TaxNumber string        `bson:"taxnumber,omitempty"`
	Address   utils.Address `bson:"address,omitempty"`
}

func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.format()
	user.Address.Format()
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
	if user.Phone == "" {
		userErrors += "phone, "
	}
	if errors := user.Address.Validate(); errors != "" {
		userErrors += errors
	}
	if userErrors != "" {
		return errors.New("There are errors in the fields: " + userErrors)
	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Phone = regexp.MustCompile(`[^0-9]+`).ReplaceAllString(user.Phone, "")
	user.TaxNumber = regexp.MustCompile(`[^0-9]+`).ReplaceAllString(user.TaxNumber, "")
}
