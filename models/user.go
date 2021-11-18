package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User : user model
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"`
	LastName  string             `bson:"lastName" json:"lastName,omitempty"`
	BirthDate time.Time          `bson:"birthDate" json:"birthDate,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Comment   string             `bson:"comment" json:"comment,omitempty"`
	Website   string             `bson:"website" json:"website,omitempty"`
}
