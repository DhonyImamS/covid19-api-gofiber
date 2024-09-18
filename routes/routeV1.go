package routes

import (
	"github.com/gofiber/fiber/v2"
	"covid19-api-gofiber/controller"
)

func RouteV1(app *fiber.App) {
	// Group API v1 routes
    apiV1 := app.Group("/v1")

	apiV1.Get("/patients", handlerGetPatients)
}

func handlerGetPatients(c *fiber.Ctx) error {
	patientController := controller.PatientController{};

	result := patientController.GetAll(c);

	return result;
}