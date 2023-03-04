package services

import (
	"context"

	"github.com/adriantoro26/meetups-core/src/meetups/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MeetupServiceDefinition struct {
	collection *mongo.Collection
	initialize bool
}

func (service *MeetupServiceDefinition) Constructor(collection *mongo.Collection) *MeetupServiceDefinition {
	service.collection = collection
	service.initialize = true
	return service
}

func (service *MeetupServiceDefinition) Create(meetup *models.Meetup) (*models.Meetup, error) {
	// Create meetup document
	response, err := service.collection.InsertOne(context.TODO(), *meetup)

	if err != nil {
		return nil, err
	}

	var entity *models.Meetup

	// Get created document
	service.collection.FindOne(context.TODO(), bson.D{{"_id", response.InsertedID}}).Decode(&entity)

	return entity, nil
}

func (service *MeetupServiceDefinition) Single(_id string) (*models.Meetup, error) {

	// Convert to mongoID
	id, _ := primitive.ObjectIDFromHex(_id)

	// Find meeetup
	var meetup *models.Meetup
	err := service.collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&meetup)

	// Return error message if meetup is not found
	if err == mongo.ErrNoDocuments {
		return nil, err
	}

	return meetup, nil
}

func (service *MeetupServiceDefinition) All() ([]models.Meetup, error) {
	// Find all meetups
	var meetups []models.Meetup
	cursor, err := service.collection.Find(context.TODO(), bson.D{})

	defer cursor.Close(context.TODO())

	err = cursor.All(context.TODO(), &meetups)

	// Return error message if meetup is not found

	if err != nil {
		return nil, err
	}

	return meetups, nil
}
