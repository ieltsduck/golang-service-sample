package main

import (
	"golang-service-sample/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Add logger middleware
	app.Use(logger.New())

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello GoFiber Project!")
	// })

	// app.Get("/about", func(c *fiber.Ctx) error {
	// 	return c.SendString("About Page")
	// })

	routes.Setup(app)

	app.Listen(":3000")
}
