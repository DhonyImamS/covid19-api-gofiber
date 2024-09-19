package controller

import (
	"covid19-api-gofiber/model"
	"covid19-api-gofiber/schema"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"log"
)

type PatientController struct {}

var validate = validator.New()

func (pc *PatientController) GetAll(c *fiber.Ctx) error {
	patientModel := model.Patient{};

	result, err := patientModel.All();
	if err != nil { 
		log.Println(err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err,
        })
	}

	if len(result) < 1 {
		dataRes := schema.ResponseEmpty{
			Message:    "Data is empty",
            StatusCode: 200,
		};
		return c.Status(fiber.StatusOK).JSON(dataRes);
	} else {
		dataRes := schema.ResponseData{
            Message:    "Get All Resource",
            StatusCode: 200,
            Data:       result,
        };
        return c.Status(fiber.StatusOK).JSON(dataRes);
	}
}

func (pc *PatientController) Store(c *fiber.Ctx) error {
	var reqData model.Patient;
	patientModel := model.Patient{};

	// Validate the struct using the validator package
    err := validate.Struct(reqData)
    if err != nil {
        // Extract and return validation errors
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

	// Parse the request body into the user struct
	if err := c.BodyParser(&reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON request",
		})
	}

	// store data into model
	err1 := patientModel.Store(reqData);
	if err1 != nil { 
		log.Println(err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err1,
        })
	}

	dataRes := schema.ResponseEmpty{
		Message:    "Data Successfully Submitted",
		StatusCode: 201,
	};
	return c.Status(fiber.StatusCreated).JSON(dataRes);
}