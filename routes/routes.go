package routes

import (
	"api_fiber_warehouse/controllers"
	modelsSess "api_fiber_warehouse/models/sessions"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

//var (
//store *session.Store
// TOKEN     string = "token"
// AUTH_KEY  string = "authenticated"
// USER_ID   string = "user_id"
// BRANCH_ID string = "id_branch"
//)

func Routes(router *fiber.App) {

	router.Use(modelsSess.NewMiddleware(), cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	//router.Use(modelsSess.NewMiddleware())

	router.Post("/check/header", func(c *fiber.Ctx) error {
		cookie := c.Get("cookie")
		//fmt.Println(cookie)
		return c.SendString(cookie)
	})

	router.Get("/", Hello)
	router.Get("/landing", controllers.Landing)
	router.Post("/login", controllers.Login)

	Coredata(router)
	Apiquote(router)
	Search(router)
}

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello fiber.")
}
