package entities

import (
	"Api-Go/pkg/entities/utils"
	"errors"
	"strings"
)

type Rent struct {
	Id                   string           `bson:"_id,omitempty"`
	UserName             string           `bson:"userName,omitempty"`
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

	if rent.UserName == "" {
		rentErrors += "UserName, "
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
		rentErrors += "DeliveryDate - delivery date must be within 30 days, "
	}
	if rentErrors != "" {
		return errors.New("There are errors in the fields: " + rentErrors)
	}
	return nil
}

func (rent *Rent) format() {
	rent.UserName = strings.TrimSpace(rent.UserName)
	rent.BookName = strings.TrimSpace(rent.BookName)
	rent.BookPublisher = strings.TrimSpace(rent.BookPublisher)
	rent.BookAuthor = strings.TrimSpace(rent.BookAuthor)
	rent.DeliveryDate.Format()
}
