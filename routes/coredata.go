package routes

import (
	"api_fiber_warehouse/controllers"

	"github.com/gofiber/fiber/v2"
)

func Coredata(router *fiber.App) {

	api := router.Group("/api/coredata")

	api.Get("/branchs/:act", controllers.Branchs)
	api.Get("/categorys/:act", controllers.Categorys)
	api.Get("/products/:act", controllers.Products)
	api.Get("/positions/:act", controllers.Positions)
	api.Get("/customers/:act", controllers.Customers)

	api.Get("/tambon", controllers.Tambon)
	api.Get("/amphure", controllers.Amphure)
	api.Get("/province", controllers.Province)
	api.Get("/tmpall", controllers.Tmpall)

	// api.Get("/branchs/:act", func(c *fiber.Ctx) error {
	// 	return c.SendString("/api/coredata/branchs/")
	// }) //
}
