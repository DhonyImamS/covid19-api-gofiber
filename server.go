package main

import (
	"github.com/gofiber/fiber/v2"
	"covid19-api-gofiber/routes"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "uas-boilerplate-go-fiber",
		Prefork:  true,
		StrictRouting: true,
		CaseSensitive: true,
	})

	app.Get("/", healthCheck)
	routes.RouteV1(app)

	app.Listen(":3000")
}

func healthCheck(c *fiber.Ctx) error {
	err := c.SendString("Hello, World!")

	if err != nil { 
		return fiber.NewError(500, "Internal Server Error")
	}

	return nil;
}
