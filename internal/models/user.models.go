package models

type User struct {
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
	Email     string `bson:"email" json:"email"`
	Username  string `bson:"username" json:"username"`
}
