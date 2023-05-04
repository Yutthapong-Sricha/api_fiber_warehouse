package sessions

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func NewMiddleware() fiber.Handler {
	return AuthMiddleware
}

func AuthMiddleware(c *fiber.Ctx) error {
	//header := c.Request()
	//Authorization := fiber.HeaderAuthorization
	//log.Println(header.String())
	// x, _ := store.Storage.Get("session_id")

	// fmt.Println(x)
	base_uri := strings.Split(c.Path(), "/")[1]
	if base_uri == "" || base_uri == "login" || base_uri == "landing" {
		return c.Next()
	}

	//fmt.Println(c.Get("Cookies"))
	fmt.Println(c.Get("Authorization"))
	//sess, err := store.Get(c)
	//fmt.Println(sess.Get("name"))
	//fmt.Println(err.Error())

	// if err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"message": "not authorized1",
	// 	})
	// }

	//coo := c.Cookies("session_id")

	// c.Request().Header.VisitAll(func(key, value []byte) {
	// 	fmt.Println("req headerKey", string(key), "value", string(value))
	// })

	// if sess.Get(USER_ID) == nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"message": "not authorized2",
	// 	})
	// }

	return c.Next()
}
