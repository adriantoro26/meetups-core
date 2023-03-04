package main

import (
	"context"
	"os"

	"github.com/adriantoro26/meetups-core/src/database"
	meetupControllers "github.com/adriantoro26/meetups-core/src/meetups/controllers"
	meetupRoutes "github.com/adriantoro26/meetups-core/src/meetups/routes"
	meetupServices "github.com/adriantoro26/meetups-core/src/meetups/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

// Global variables
var (
	mongoClient             *mongo.Client
	meetupModel             *mongo.Collection
	meetupController        meetupControllers.MeetupController
	meetupRoute             meetupRoutes.MeetupRoutes
	meetupService           meetupServices.MeetupService
	meetupServiceDefinition meetupServices.MeetupServiceDefinition
)

func init() {
	// Initialize package dependencies

	// Get env variables
	mongoUri, found := os.LookupEnv("DB_MONGO_URI")

	if found == false {
		mongoUri = "mongodb://localhost:27017"
	}

	// Connect to MongoDB database
	mongoClient = database.MongoDBConnect(mongoUri)

	// Get Meetup collection
	meetupModel = database.GetMongoCollection(mongoClient, "project", "meetup")

	// Initialize controllers
	meetupService = meetupServiceDefinition.Constructor(meetupModel)
	meetupController.Service = meetupService
	meetupRoute.Constructor(&meetupController)
}

// Application entry point
func main() {

	// Get environment variables
	apiKey := os.Getenv("API_KEY")

	defer mongoClient.Disconnect(context.TODO())

	// Instantiate echo framework
	e := echo.New()

	// Register Middlewares
	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == apiKey, nil
	}))

	// Register routes and handlers
	meetupRoute.RegisterRoutes(e)

	// Start server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
