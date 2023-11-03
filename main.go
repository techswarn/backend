package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techswarn/backend/routes"

	"fmt"
	"os"

	"github.com/techswarn/backend/database"
	"github.com/techswarn/backend/utils"
)

// define the default port of the application
const DEFAULT_PORT = "3000"

// NewFiberApp returns fiber application
func NewFiberApp() *fiber.App {
	// create a new fiber application
	var app *fiber.App = fiber.New()

	// define the routes
	routes.SetupRoutes(app)

	// return the application
	return app
}

func main() {
	// create a new fiber application
	var app *fiber.App = NewFiberApp()

	// connect to the database
	database.InitDatabase(utils.GetValue("DB_NAME"))

	// get the application port from the defined PORT variable
	var PORT string = os.Getenv("PORT")

	// if the PORT variable is not assigned
	// use the default port
	if PORT == "" {
		PORT = DEFAULT_PORT
	}

	// start the application
	app.Listen(fmt.Sprintf(":%s", PORT))
}
