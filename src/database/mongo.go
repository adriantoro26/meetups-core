package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// description: Open connection to MongoDB database
// param: uri - MongoDB URI string
func MongoDBConnect(uri string) *mongo.Client {

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	return client
}

// description: Get acess to mongo collection
// param: database - Database name
// param: colleciton - Collection name
func GetMongoCollection(client *mongo.Client, database string, collection string) *mongo.Collection {

	// Get collection
	model := client.Database(database).Collection(collection)

	return model
}
