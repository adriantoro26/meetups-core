package controllers

import (
	"fmt"
	"net/http"

	"github.com/adriantoro26/meetups-core/src/meetups/models"
	"github.com/adriantoro26/meetups-core/src/meetups/services"

	"github.com/labstack/echo/v4"
)

// Custom types
type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type MeetupController struct {
	service services.MeetupService
}

func New(service services.MeetupService) *MeetupController {
	return &MeetupController{service}
}

// description: Creates a new meetup
func (controller MeetupController) Create(c echo.Context) error {

	meetup := &models.Meetup{}

	// Get request body
	if err := c.Bind(meetup); err != nil {
		return err
	}

	// Create meetup document
	meetup, err := controller.service.Create(meetup)

	if err != nil {
		response := &Response{"500", "Internal Server Error"}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusCreated, *meetup)
}

// description: Get single meetup
func (controller MeetupController) Single(c echo.Context) error {

	// Get param id
	_id := c.Param("_id")

	// Find meeetup
	meetup, err := controller.service.Single(_id)

	// Return error message if meetup is not found
	if err != nil {
		response := &Response{"404", fmt.Sprintf("No meetup found with given id: %s\n", _id)}
		return c.JSON(http.StatusNotFound, response)
	}

	return c.JSON(http.StatusOK, *meetup)
}

// description: Get all meetups
func (controller MeetupController) All(c echo.Context) error {

	// Find all meetups
	meetups, err := controller.service.All()

	// Return error message if meetup is not found

	if err != nil {
		response := &Response{"500", "Internal Server Error"}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, meetups)
}
