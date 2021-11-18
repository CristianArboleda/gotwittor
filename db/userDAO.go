package db

import (
	"context"
	"time"

	"github.com/CristianArboleda/gotwittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// SaveUser : save user in the DB
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

// UpdateUser : save user in the DB
func UpdateUser(us models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConnection.Database("gotwitor")
	collection := db.Collection("user")

	registro := make(map[string]interface{})
	if len(us.Name) > 0 {
		registro["name"] = us.Name
	}
	if len(us.LastName) > 0 {
		registro["lastName"] = us.LastName
	}
	if len(us.Comment) > 0 {
		registro["comment"] = us.Comment
	}
	if len(us.Website) > 0 {
		registro["website"] = us.Website
	}
	if len(us.Avatar) > 0 {
		registro["avatar"] = us.Avatar
	}
	if len(us.Banner) > 0 {
		registro["banner"] = us.Banner
	}
	if len(us.Banner) > 0 {
		registro["banner"] = us.Banner
	}
	if !us.BirthDate.IsZero() {
		registro["birthDate"] = us.BirthDate
	}

	if len(us.Password) > 0 {
		registro["password"], _ = EncryptPass(us.Password)
	}

	updateString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{
		"_id": bson.M{
			"$eq": objID,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}

// FindUserByEmail : find if exist a user by a email
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

// FindUserById : find if exist a user by an ID
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

// CheckLogin : check if login params are valid
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
