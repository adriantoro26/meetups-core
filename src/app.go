package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

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
