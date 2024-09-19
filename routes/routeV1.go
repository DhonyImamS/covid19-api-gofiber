package routes

import (
	"github.com/gofiber/fiber/v2"
	"covid19-api-gofiber/controller"
)

func RouteV1(app *fiber.App) {
	patientController := controller.PatientController{};

	// Group API v1 routes
    apiV1 := app.Group("/v1")

	apiV1.Get("/patients", patientController.GetAll)
	apiV1.Post("/patients/submit", patientController.Store)
}
