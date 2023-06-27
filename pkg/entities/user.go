package entities

import (
	"Api-Go/pkg/entities/utils"
	"errors"
	"regexp"
	"strings"

	"github.com/klassmann/cpfcnpj"
)

type User struct {
	Id        string        `bson:"_id,omitempty"`
	Name      string        `bson:"name,omitempty"`
	FullName  string        `bson:"fullName,omitempty"`
	Email     string        `bson:"email,omitempty"`
	Phone     string        `bson:"phone,omitempty"`
	TaxNumber string        `bson:"taxNumber,omitempty"`
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

	if user.Name == "" {
		userErrors += "name, "
	}
	if user.FullName == "" || !utils.ValidateName(user.FullName) {
		userErrors += "fullname, "
	}
	if user.Email == "" || !utils.ValidateEmail(user.Email) {
		userErrors += "email, "
	}
	if user.Phone == "" {
		userErrors += "phone, "
	}
	if user.TaxNumber == "" || !cpfcnpj.ValidateCPF(user.TaxNumber) || !user.ValidateCPF() {
		userErrors += "taxnumber, "
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
	user.FullName = strings.TrimSpace(user.FullName)
	user.Email = strings.TrimSpace(user.Email)
	user.Phone = strings.TrimSpace(user.Phone)
	user.TaxNumber = strings.TrimSpace(user.TaxNumber)

	user.Phone = regexp.MustCompile(`[^0-9]+`).ReplaceAllString(user.Phone, "")
	user.TaxNumber = regexp.MustCompile(`[^0-9]+`).ReplaceAllString(user.TaxNumber, "")
}

func (user *User) ValidateCPF() bool {
	var invalidCpfs = []string{"00000000000", "11111111111", "22222222222", "33333333333", "44444444444", "55555555555", "66666666666", "77777777777", "88888888888", "99999999999"}
	for _, cpf := range invalidCpfs {
		if user.TaxNumber == cpf {
			return false
		}
	}

	return true
}
