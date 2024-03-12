package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/techswarn/backend/routes"
	"github.com/techswarn/backend/database"
	"github.com/techswarn/backend/utils"

	"fmt"
	"os"


)

// define the default port of the application
const DEFAULT_PORT = "3000"

// NewFiberApp returns fiber application
func NewFiberApp() *fiber.App {
	// create a new fiber application
	var app *fiber.App = fiber.New()

	app.Use(cors.New())

	// app.Use(cors.New(cors.Config{
	// 	AllowOriginsFunc: func(origin string) bool {
	// 		return os.Getenv("ENVIRONMENT") == "development"
	// 	},
	// }))
	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
        AllowOrigins:     "*",
        AllowCredentials: true,
        AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))
		
	//Loging middleware
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	// define the routes
	routes.SetupRoutes(app)

	// return the application
	return app
}

func main() {
	fmt.Println("Go backend")
	// create a new fiber application
	var app *fiber.App = NewFiberApp()
	// Initialize default config
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
