package main

import (
	"context"
	"fmt"

	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Meetup schema
type Meetup struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title" binding:"required"`
	Description string             `json:"description" bson:"description" binding:"required"`
	Image       string             `json:"image,omitempty" bson:"image" binding:"required"`
	Address     string             `json:"address" bson:"address" binding:"required"`
}
// Global variables
var meetupModel *mongo.Collection

// description: Open connection to MongoDB database
func mongoDBConnect(uri string, database string, collection string) *mongo.Collection {

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

	model := client.Database(database).Collection(collection)

	return model
}

// Route handlers

// description: Creates a new meetup
func createMeetup(c echo.Context) error {
	fmt.Println("Meetup created")
	return nil
}

// description: Get single meetup
func getSingleMeetup(c echo.Context) error {
	fmt.Println("Single meetup returned")
	return nil
}

// description: Get all meetups
func getAllMeetup(c echo.Context) error {
	fmt.Println("All meetups returned")
	return nil
}

// Application entry point
func main() {

	// Instantiate echo framework
	e := echo.New()

	// Register routes and handlers
	e.GET("/meetups", getAllMeetup)
	e.GET("/meetups/:_id", getSingleMeetup)
	e.POST("/meetups", createMeetup)

	// Start server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
