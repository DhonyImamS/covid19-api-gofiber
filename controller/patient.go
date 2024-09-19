package controller

import (
	"covid19-api-gofiber/model"
	"covid19-api-gofiber/types"
	"github.com/gofiber/fiber/v2"
)

type PatientController struct {}

func (pc *PatientController) GetAll(c *fiber.Ctx) error {
	patientModel := model.Patient{};

	result := patientModel.All();

	if len(result) < 1 {
		dataRes := types.ResponseEmpty{
			Message:    "Data is empty",
            StatusCode: 200,
		};
		return c.Status(fiber.StatusOK).JSON(dataRes);
	} else {
		dataRes := types.ResponseData{
            Message:    "Get All Resource",
            StatusCode: 200,
            Data:       result,
        };
        return c.Status(fiber.StatusOK).JSON(dataRes);
	}
}