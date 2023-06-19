package models

import (
	"errors"
	"strings"
)

type Book struct {
	ID          string `bson:"id,omitempty"`
	Name        string `bson:"name,omitempty"`
	Author      string `bson:"author,omitempty"`
	Publisher   string `bson:"publisher,omitempty"`
	ReleaseDate string `bson:"releaseDate,omitempty"`
	Stock       uint   `bson:"stock,omitempty"`
	Rented      uint   `bson:"rented"`
}

func (book *Book) Prepare() error {
	if err := book.validate(); err != nil {
		return err
	}

	book.format()
	return nil
}
func (book *Book) validate() error {
	bookErrors := ""

	if book.Name == "" {
		bookErrors += "name, "
	}
	if book.Author == "" {
		bookErrors += "Author, "
	}
	if book.Publisher == "" {
		bookErrors += "Publisher, "
	}
	if book.ReleaseDate == "" {
		bookErrors += "ReleaseDate, "
	}
	if book.Stock == 0 {
		bookErrors += "Stock, "
	}
	if bookErrors != "" {
		return errors.New("There are errors in the fields: " + bookErrors)
	}
	return nil
}

func (book *Book) format() {
	book.Name = strings.TrimSpace(book.Name)
	book.Author = strings.TrimSpace(book.Author)
	book.Publisher = strings.TrimSpace(book.Publisher)
	book.ReleaseDate = strings.TrimSpace(book.ReleaseDate)
}
