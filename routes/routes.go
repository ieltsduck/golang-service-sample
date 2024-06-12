package routes

import (
	"golang-service-sample/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Get("/about", handlers.About)
}
