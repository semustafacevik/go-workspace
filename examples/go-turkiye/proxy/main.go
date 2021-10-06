package main

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/semustafacevik/go-workspace/go-turkiye/proxy"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to go-workspace | proxy")
	})

	app.Get("/:key/*", ProxyHandler)

	app.Listen(":3000")
}
