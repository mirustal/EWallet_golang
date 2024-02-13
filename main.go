package main

import (
	"ewallet/app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "",
		AllowHeaders:  "Origin, Content-Type, Accept",
	}))

	// app.Use(accessMiddleware)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})

	routes.Init(app)

	app.Listen(":8080")

}

// func accessMiddleware(c *fiber.Ctx) error {
// 	accessToken := c.Params("acces_token")
// 	if !hasAccess(accesToken) {
// 		return c.SendStatus(fiber.StatusUnauthorized)
// 	}
// 	return c.Next()
// }