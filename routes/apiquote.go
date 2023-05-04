package routes

import (
	"api_fiber_warehouse/controllers"

	"github.com/gofiber/fiber/v2"
)

func Apiquote(router *fiber.App) {
	//func Coredata(routes *fiber.Route) {
	api := router.Group("/api/quotation")
	api.Get("/promotions/:act", controllers.Promotions)

	api.Post("/createtmpcore", controllers.CreateTmpCore)
	api.Post("/listtmpcore", controllers.ListTmpCore)
}
