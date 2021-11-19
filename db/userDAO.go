package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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

	record := make(map[string]interface{})
	if len(us.Name) > 0 {
		record["name"] = us.Name
	}
	if len(us.LastName) > 0 {
		record["lastName"] = us.LastName
	}
	if len(us.Comment) > 0 {
		record["comment"] = us.Comment
	}
	if len(us.Website) > 0 {
		record["website"] = us.Website
	}
	if len(us.Avatar) > 0 {
		record["avatar"] = us.Avatar
	}
	if len(us.Banner) > 0 {
		record["banner"] = us.Banner
	}
	if len(us.Banner) > 0 {
		record["banner"] = us.Banner
	}
	if !us.BirthDate.IsZero() {
		record["birthDate"] = us.BirthDate
	}

	if len(us.Password) > 0 {
		record["password"], _ = EncryptPass(us.Password)
	}

	updateString := bson.M{
		"$set": record,
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

// FindUsersByFilters : Find users by filters
func FindUsersByFilters(ID string, page int64, search string, relationType string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("gotwitor")
	collection := db.Collection("user")

	var result []*models.User
	opts := options.Find()
	opts.SetSkip((page - 1) * 20)
	opts.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}
	records, err := collection.Find(ctx, query, opts)

	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	var add bool

	for records.Next(context.TODO()) {
		var usu models.User
		err = records.Decode(&usu)
		if err != nil {
			log.Fatal(err.Error())
			return result, false
		}
		if ID == usu.ID.Hex() {
			continue
		}
		rel, _ := FindRelationByUserAndRelationUser(ID, usu.ID.Hex())
		if ID == rel.RelationUserID {
			continue
		}
		add = false
		if relationType == "NO_FOLLOW" && len(rel.UserID) < 1 {
			add = true
		} else if relationType == "FOLLOW" && len(rel.UserID) > 0 {
			add = true
		}

		if add {
			usu.Password = ""
			result = append(result, &usu)
		}
	}
	records.Close(context.TODO())
	return result, true
}
