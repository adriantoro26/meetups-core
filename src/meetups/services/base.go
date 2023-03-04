package services

import "github.com/adriantoro26/meetups-core/src/meetups/models"

type MeetupService interface {
	Create(*models.Meetup) (*models.Meetup, error)
	Single(string) (*models.Meetup, error)
	All() ([]models.Meetup, error)
}
