package db

import (
	"context"
	"github.com/CristianArboleda/gotwittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// SaveRelation : Save relation between two users
func SaveRelation(re models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("gotwitor")
	collection := db.Collection("relation")

	_, err := collection.InsertOne(ctx, re)
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeleteRelation : Delete a relation between two users
func DeleteRelation(re models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("gotwitor")
	collection := db.Collection("relation")

	_, err := collection.DeleteOne(ctx, re)
	if err != nil {
		return false, err
	}
	return true, nil
}

// FindRelationByUserAndRelationUser : find if exist a relation between two users
func FindRelationByUserAndRelationUser(userID string, relationUserID string) (models.Relation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("gotwitor")
	collection := db.Collection("relation")

	condition := bson.M{
		"userid":         userID,
		"relationuserid": relationUserID,
	}

	var result models.Relation

	err := collection.FindOne(ctx, condition).Decode(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}
