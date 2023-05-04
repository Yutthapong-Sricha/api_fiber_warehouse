package controllers

import (
	modelsCore "api_fiber_warehouse/models/apiCoreData"

	"github.com/gofiber/fiber/v2"
)

func Branchs(c *fiber.Ctx) error {
	act := c.Params("act")
	branchs := modelsCore.ListBranch(act)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": branchs})
}

func Categorys(c *fiber.Ctx) error {
	act := c.Params("act")
	categorys := modelsCore.ListCategory(act)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": categorys})
}

func Products(c *fiber.Ctx) error {
	act := c.Params("act")
	products := modelsCore.ListProduct(act)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": products})
}

func Positions(c *fiber.Ctx) error {
	act := c.Params("act")
	positions := modelsCore.ListPosition(act)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": positions})
}

func Customers(c *fiber.Ctx) error {
	act := c.Params("act")
	customers := modelsCore.ListCustomer(act)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": customers})
}

func Tambon(c *fiber.Ctx) error {
	tambons := modelsCore.Listtambon(c)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": tambons})
}

func Amphure(c *fiber.Ctx) error {
	amphures := modelsCore.Listamphure(c)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": amphures})
}

func Province(c *fiber.Ctx) error {
	provinces := modelsCore.Listprovince(c)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": provinces})
}

func Tmpall(c *fiber.Ctx) error {
	tambons := modelsCore.Listtambon(c)
	amphures := modelsCore.Listamphure(c)
	provinces := modelsCore.Listprovince(c)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"tambons": tambons, "amphures": amphures, "provinces": provinces})
}
