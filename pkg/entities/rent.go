package entities

import (
	"Api-Go/pkg/entities/utils"
	"errors"
	"strings"
)

type Rent struct {
	Id                   string           `bson:"_id,omitempty"`
	ClientName           string           `bson:"clientName,omitempty"`
	BookName             string           `bson:"bookName,omitempty"`
	BookPublisher        string           `bson:"bookPublisher,omitempty"`
	BookAuthor           string           `bson:"bookAuthor,omitempty"`
	Delivered            bool             `bson:"delivered"`
	DeliveryDate         utils.CustomTime `bson:"deliveryDate,omitempty"`
	ExpectedDeliveryDate utils.CustomTime `bson:"expectedDeliveryDate,omitempty"`
}

func (rent *Rent) Prepare() error {
	if err := rent.validate(); err != nil {
		return err
	}

	rent.format()
	return nil
}
func (rent *Rent) validate() error {
	rentErrors := ""

	if rent.ClientName == "" {
		rentErrors += "ClientName, "
	}
	if rent.BookName == "" {
		rentErrors += "BookName, "
	}
	if rent.BookPublisher == "" {
		rentErrors += "BookPublisher, "
	}
	if rent.BookAuthor == "" {
		rentErrors += "BookAuthor, "
	}
	if !utils.ValidateDeliveryDate(rent.ExpectedDeliveryDate) {
		rentErrors += "DeliveryDate - the delivery date must be within 30 days from today, "
	}
	if rentErrors != "" {
		return errors.New("There are errors in the fields: " + rentErrors)
	}
	return nil
}

func (rent *Rent) format() {
	rent.ClientName = strings.TrimSpace(rent.ClientName)
	rent.BookName = strings.TrimSpace(rent.BookName)
	rent.BookPublisher = strings.TrimSpace(rent.BookPublisher)
	rent.BookAuthor = strings.TrimSpace(rent.BookAuthor)
	rent.DeliveryDate.Format()
}
