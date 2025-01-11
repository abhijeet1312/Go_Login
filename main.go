package main

import (
	"GoFiber-Login/app/routes"
	"GoFiber-Login/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Serve static files from the "public" directory
	app.Static("/", "./public")

	// Connect to the MongoDB database
	config.ConnectDB()

	// Register user routes
	routes.UserRoutes(app)

	// Start the server on port 3000
	app.Listen(":3000")
}
