package search

import (
	modelsSearch "api_fiber_warehouse/models/apiSearch"

	"github.com/gofiber/fiber/v2"
)

func Searchcust(c *fiber.Ctx) error {
	listSearch := modelsSearch.Cust(c)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Success Search", "data": listSearch})
}

func Searchprod(c *fiber.Ctx) error {
	listSearch := modelsSearch.Prod(c)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Success Search", "data": listSearch})
}
