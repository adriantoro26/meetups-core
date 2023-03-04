package meetups

import (
	"github.com/adriantoro26/meetups-core/src/database"
	meetupControllers "github.com/adriantoro26/meetups-core/src/meetups/controllers"
	meetupRoutes "github.com/adriantoro26/meetups-core/src/meetups/routes"
	meetupServices "github.com/adriantoro26/meetups-core/src/meetups/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(client *mongo.Client) *meetupRoutes.MeetupRoutes {

	var meetupService meetupServices.MeetupService

	// Get Meetup collection
	meetupModel := database.GetMongoCollection(client, "project", "meetup")

	// Initialize controllers
	meetupService = meetupServices.New(meetupModel)
	meetupController := meetupControllers.New(meetupService)
	meetupRoute := meetupRoutes.New(meetupController)

	return meetupRoute
}
