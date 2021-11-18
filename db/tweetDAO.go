package db

import (
	"context"
	"log"
	"time"

	"github.com/CristianArboleda/gotwittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SaveTweet : save a tweet in the DB
func SaveTweet(tw models.Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("gotwitor")
	collection := db.Collection("tweet")

	tweet := bson.M{
		"userid":     tw.UserID,
		"message":    tw.Message,
		"createdate": tw.CreateDate,
	}

	result, err := collection.InsertOne(ctx, tweet)

	if err != nil {
		return "", false, err
	}
	objI, _ := result.InsertedID.(primitive.ObjectID)

	return objI.Hex(), true, nil
}

// DeleteTweet : save a tweet in the DB
func DeleteTweet(ID string, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("gotwitor")
	collection := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id":    objID,
		"userid": userID,
	}
	_, err := collection.DeleteOne(ctx, condition)

	return err
}

// FindTweetsByUserID : find all tweets of a user
func FindTweetsByUserID(userID string, page int64) ([]*models.Tweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("gotwitor")
	collection := db.Collection("tweet")

	var result []*models.Tweet
	condition := bson.M{"userid": userID}
	opts := options.Find()
	opts.SetLimit(20)
	opts.SetSort(bson.D{{Key: "createdate", Value: -1}})
	opts.SetSkip((page - 1) * 20)

	records, err := collection.Find(ctx, condition, opts)

	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}
	err = records.All(context.TODO(), &result)
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}
	/* old implementation example:
	for records.Next(context.TODO()){
		var tweet models.Tweet
		err = records.Decode(&tweet)
		if err != nil {
			log.Fatal(err.Error())
			return result, false
		}
		result = append(result, &tweet)
	}
	*/

	return result, true
}
