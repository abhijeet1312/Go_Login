package routes

import (
	"GoFiber-Login/app/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	// Serve the home page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./public/home.html")
	})

	// Serve the register page
	app.Get("/register", controller.RegisterPage)
	app.Post("/register", controller.Register)

	// Serve the login page
	app.Get("/login", controller.LoginPage)
	app.Post("/login", controller.Login)
}
