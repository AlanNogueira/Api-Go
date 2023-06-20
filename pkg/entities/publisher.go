package entities

import (
	"Api-Go/pkg/entities/utils"
	"errors"
	"regexp"
	"strings"
)

type Publisher struct {
	Id    string `bson:"_id,omitempty"`
	Name  string `bson:"name,omitempty"`
	Email string `bson:"email,omitempty"`
	Phone string `bson:"phone,omitempty"`
	Site  string `bson:"site,omitempty"`
}

func (publisher *Publisher) Prepare() error {
	if err := publisher.validate(); err != nil {
		return err
	}

	publisher.format()
	return nil
}
func (publisher *Publisher) validate() error {
	publisherErrors := ""

	if publisher.Name == "" || !utils.ValidateName(publisher.Name) {
		publisherErrors += "name, "
	}
	if publisher.Email == "" || !utils.ValidateEmail(publisher.Email) {
		publisherErrors += "email, "
	}
	if publisher.Phone == "" {
		publisherErrors += "phone, "
	}
	if publisherErrors != "" {
		return errors.New("There are errors in the fields: " + publisherErrors)
	}
	return nil
}

func (publisher *Publisher) format() {
	publisher.Name = strings.TrimSpace(publisher.Name)
	publisher.Email = strings.TrimSpace(publisher.Email)
	publisher.Phone = regexp.MustCompile(`[^0-9]+`).ReplaceAllString(publisher.Phone, "")
}
