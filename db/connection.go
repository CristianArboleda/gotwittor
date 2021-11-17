package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoConnection: Mongo connection object */
var MongoConnection = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://cristian:Cr1st14n@gotwittor.dpz3j.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/* ConnectDB: Create the connection to Mongo DB*/
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to Mongo:  %v", err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Error in the ping to Mongo: %v", err.Error())
		return client
	}
	log.Println("Mongo conection successfully")
	return client
}

/*CheckConnection: check if the connection is active*/
func CheckConnection() bool {
	err := MongoConnection.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Error in the ping to Mongo: %v", err.Error())
		return false
	}
	return true
}
