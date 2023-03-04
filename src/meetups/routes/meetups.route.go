package routes

import (
	"github.com/adriantoro26/meetups-core/src/meetups/controllers"
	"github.com/labstack/echo/v4"
)

type MeetupRoutes struct {
	meetupController *controllers.MeetupController
}

func New(meetupController *controllers.MeetupController) *MeetupRoutes {
	return &MeetupRoutes{meetupController}
}

func (p *MeetupRoutes) RegisterRoutes(e *echo.Echo) {
	// Register routes and handlers
	e.GET("/meetups", p.meetupController.All)
	e.GET("/meetups/:_id", p.meetupController.Single)
	e.POST("/meetups", p.meetupController.Create)
}
