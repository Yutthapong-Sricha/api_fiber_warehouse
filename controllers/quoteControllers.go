package controllers

import (
	modelsQuote "api_fiber_warehouse/models/apiQuote"

	"github.com/gofiber/fiber/v2"
)

func Promotions(c *fiber.Ctx) error {
	act := c.Params("act")
	promotions := modelsQuote.ListPromotion(act)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": promotions})
}

func CreateTmpCore(c *fiber.Ctx) error {
	tmpdoccore, err := modelsQuote.Createtmpcore(c)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "error", "message": "Create Error : " + err.Error(), "data": tmpdoccore,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Success Created", "data": tmpdoccore})
}

func ListTmpCore(c *fiber.Ctx) error {
	listtmpcore, err := modelsQuote.Listtmpcore(c)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "error", "message": "Create Error : " + err.Error(), "data": listtmpcore,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Success list", "data": listtmpcore})
}
