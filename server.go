package main

import (
	"api_fiber_warehouse/initializers"
	setup "api_fiber_warehouse/routes"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {

	app := fiber.New()
	//app.Use(cors.New())
	// app.Get("/landing", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{"status": "success", "branchs": "hello world"})
	// })

	app.Use(favicon.New())
	setup.Routes(app)
	//app.Listen(":3008")
	app.Listen(":" + os.Getenv("PORT"))
}
