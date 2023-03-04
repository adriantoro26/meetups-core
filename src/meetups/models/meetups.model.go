package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Meetup schema
type Meetup struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title" binding:"required"`
	Description string             `json:"description" bson:"description" binding:"required"`
	Image       string             `json:"image,omitempty" bson:"image" binding:"required"`
	Address     string             `json:"address" bson:"address" binding:"required"`
}
