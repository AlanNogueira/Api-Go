package entities

import (
	"Api-Go/pkg/entities/utils"
	"errors"
	"regexp"
	"strings"

	"github.com/klassmann/cpfcnpj"
)

type Client struct {
	Id        string        `bson:"_id,omitempty"`
	Name      string        `bson:"clientName,omitempty"`
	FullName  string        `bson:"fullName,omitempty"`
	Email     string        `bson:"email,omitempty"`
	Phone     string        `bson:"phone,omitempty"`
	TaxNumber string        `bson:"taxNumber,omitempty"`
	Address   utils.Address `bson:"address,omitempty"`
}

func (client *Client) Prepare() error {
	if err := client.validate(); err != nil {
		return err
	}

	client.format()
	client.Address.Format()
	return nil
}
func (client *Client) validate() error {
	clientErrors := ""

	if client.Name == "" {
		clientErrors += "name, "
	}
	if client.FullName == "" || !utils.ValidateName(client.FullName) {
		clientErrors += "fullname, "
	}
	if client.Email == "" || !utils.ValidateEmail(client.Email) {
		clientErrors += "email, "
	}
	if client.Phone == "" {
		clientErrors += "phone, "
	}
	if client.TaxNumber == "" || !cpfcnpj.ValidateCPF(client.TaxNumber) || !client.ValidateCPF() {
		clientErrors += "taxnumber, "
	}
	if errors := client.Address.Validate(); errors != "" {
		clientErrors += errors
	}
	if clientErrors != "" {
		return errors.New("There are errors in the fields: " + clientErrors)
	}
	return nil
}

func (client *Client) format() {
	client.Name = strings.TrimSpace(client.Name)
	client.FullName = strings.TrimSpace(client.FullName)
	client.Email = strings.TrimSpace(client.Email)
	client.Phone = strings.TrimSpace(client.Phone)
	client.TaxNumber = strings.TrimSpace(client.TaxNumber)

	client.Phone = regexp.MustCompile(`[^0-9]+`).ReplaceAllString(client.Phone, "")
	client.TaxNumber = regexp.MustCompile(`[^0-9]+`).ReplaceAllString(client.TaxNumber, "")
}

func (client *Client) ValidateCPF() bool {
	var invalidCpfs = []string{"00000000000", "11111111111", "22222222222", "33333333333", "44444444444", "55555555555", "66666666666", "77777777777", "88888888888", "99999999999"}
	for _, cpf := range invalidCpfs {
		if client.TaxNumber == cpf {
			return false
		}
	}

	return true
}
