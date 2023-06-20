package entities

import (
	"strings"
)

type Address struct {
	Street       string `bson:"street,omitempty"`
	Number       uint   `bson:"number,omitempty"`
	Neighborhood string `bson:"neighborhood,omitempty"`
	City         string `bson:"city,omitempty"`
	State        string `bson:"state,omitempty"`
}

func (a *Address) validate() string {
	addressErrors := ""
	address := Address{}
	if address == *a {
		return "address is required, "
	}
	if a.Street == "" {
		addressErrors += "street, "
	}
	if a.Number == 0 {
		addressErrors += "number, "
	}
	if a.Neighborhood == "" {
		addressErrors += "neighborhood, "
	}
	if a.City == "" {
		addressErrors += "city, "
	}
	if a.State == "" {
		addressErrors += "state, "
	}

	return addressErrors
}

func (a *Address) format() {
	a.Street = strings.TrimSpace(a.Street)
	a.Neighborhood = strings.TrimSpace(a.Neighborhood)
	a.City = strings.TrimSpace(a.City)
	a.State = strings.TrimSpace(a.State)
}
