package db

import (
	"context"
	"time"

	"github.com/CristianArboleda/gotwittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

/*SaveUser: save user in the DB*/
func SaveUser(us models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("gotwitor")
	collection := db.Collection("user")

	us.Password, _ = EncryptPass(us.Password)

	result, err := collection.InsertOne(ctx, us)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}

/*FindUserByEmail: find if exist a user by a email*/
func FindUserByEmail(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("gotwitor")
	collection := db.Collection("user")

	condition := bson.M{"email": email}

	var result models.User

	err := collection.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}

/*FindUserById: find if exist a user by an ID*/
func FindUserById(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("gotwitor")
	collection := db.Collection("user")
	var result models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{"_id": objID}

	err := collection.FindOne(ctx, condition).Decode(&result)
	result.Password = ""

	if err != nil {
		return result, err
	}
	return result, nil
}

/*CheckLogin: check if login params are valid*/
func CheckLogin(email, pass string) (models.User, bool) {
	us, exist, _ := FindUserByEmail(email)
	if !exist {
		return us, false
	}
	passwordBytes := []byte(pass)
	passwordDB := []byte(us.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return us, false
	}
	return us, true
}
