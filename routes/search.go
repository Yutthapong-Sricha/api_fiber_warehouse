package routes

import (
	ctrlSearch "api_fiber_warehouse/controllers/search"

	"github.com/gofiber/fiber/v2"
)

func Search(router *fiber.App) {

	search := router.Group("/api/search")

	search.Post("/cust", ctrlSearch.Searchcust)
	search.Post("/prod", ctrlSearch.Searchprod)

}
