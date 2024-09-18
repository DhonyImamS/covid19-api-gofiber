package controller

import (
	"covid19-api-gofiber/model"
	"github.com/gofiber/fiber/v2"
)

type PatientController struct {}

func (pc *PatientController) GetAll(c *fiber.Ctx) error {
	patientModel := model.Patient{};

	result := patientModel.All();

	return c.JSON(result);
}