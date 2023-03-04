package main

import (
	"context"
	"os"

	"github.com/adriantoro26/meetups-core/src/database"
	"github.com/adriantoro26/meetups-core/src/meetups"
	meetupRoutes "github.com/adriantoro26/meetups-core/src/meetups/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

// Global variables
var (
	mongoClient *mongo.Client
	meetupRoute *meetupRoutes.MeetupRoutes
)

func init() {
	// Initialize modules

	// Get env variables
	mongoUri, found := os.LookupEnv("DB_MONGO_URI")

	if found == false {
		mongoUri = "mongodb://localhost:27017"
	}

	// Connect to MongoDB database
	mongoClient = database.MongoDBConnect(mongoUri)

	// Initialize meetups module
	meetupRoute = meetups.Init(mongoClient)
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
