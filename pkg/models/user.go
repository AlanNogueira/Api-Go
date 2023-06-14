package models

type User struct {
	ID        string `bson:"id,omitempty"`
	Name      string `bson:"name,omitempty"`
	Age       uint   `bson:"age,omitempty"`
	Email     string `bson:"email,omitempty"`
	Password  string `bson:"password,omitempty"`
	TaxNumber string `bson:"taxnumber,omitempty"`
}
