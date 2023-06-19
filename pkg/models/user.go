package models

import (
	"errors"
	"net/mail"
	"regexp"
	"strings"
)

type User struct {
	ID        string  `bson:"id,omitempty"`
	Name      string  `bson:"name,omitempty"`
	Email     string  `bson:"email,omitempty"`
	Phone     string  `bson:"phone,omitempty"`
	TaxNumber string  `bson:"taxnumber,omitempty"`
	Address   Address `bson:"address,omitempty"`
}

func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.format()
	user.Address.format()
	return nil
}
func (user *User) validate() error {
	userErrors := ""

	if user.Name == "" || !validateName(user.Name) {
		userErrors += "name, "
	}
	if user.Email == "" || !validateEmail(user.Email) {
		userErrors += "email, "
	}
	if user.Phone == "" {
		userErrors += "phone, "
	}
	if errors := user.Address.validate(); errors != "" {
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

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validateName(name string) bool {
	match, _ := regexp.MatchString(`^((\b[A-zÀ-ú']{2,40}\b)\s*){2,}$`, name)
	return match
}
