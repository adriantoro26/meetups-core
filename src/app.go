package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/joho/godotenv"
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

// Custom types
type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
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

	meetup := &Meetup{}

	// Get request body
	if err := c.Bind(meetup); err != nil {
		return err
	}

	// Create meetup document
	response, err := meetupModel.InsertOne(context.TODO(), meetup)

	if err != nil {
		response := &Response{"500", "Internal Server Error"}
		return c.JSON(http.StatusInternalServerError, response)
	}

	// Get created document
	meetupModel.FindOne(context.TODO(), bson.D{{"_id", response.InsertedID}}).Decode(&meetup)

	return c.JSON(http.StatusCreated, meetup)
}

// description: Get single meetup
func getSingleMeetup(c echo.Context) error {
	fmt.Println("Single meetup returned")

	// Get param id
	id := c.Param("_id")

	// Convert to mongoID
	_id, _ := primitive.ObjectIDFromHex(id)

	// Find meeetup
	var meetup Meetup
	err := meetupModel.FindOne(context.TODO(), bson.D{{"_id", _id}}).Decode(&meetup)

	// Return error message if meetup is not found
	if err == mongo.ErrNoDocuments {
		response := &Response{"404", fmt.Sprintf("No meetup found with given id: %s\n", id)}
		return c.JSON(http.StatusNotFound, response)
	}

	return c.JSON(http.StatusOK, meetup)
}

// description: Get all meetups
func getAllMeetup(c echo.Context) error {
	fmt.Println("All meetups returned")

	// Find all meetups
	var meetups []Meetup
	cursor, err := meetupModel.Find(context.TODO(), bson.D{})

	err = cursor.All(context.TODO(), &meetups)

	// Return error message if meetup is not found

	if err != nil {
		response := &Response{"500", "Internal Server Error"}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, meetups)
}

// Application entry point
func main() {

	// Load environment variables
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	mongoUri, found := os.LookupEnv("DB_MONGO_URI")

	if found == false {
		mongoUri = "mongodb://localhost:27017"
	}

	// Instantiate echo framework
	e := echo.New()

	// Register routes and handlers
	e.GET("/meetups", getAllMeetup)
	e.GET("/meetups/:_id", getSingleMeetup)
	e.POST("/meetups", createMeetup)

	// Connect to DB
	meetupModel = mongoDBConnect(mongoUri, "project", "meetup")

	// Start server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
