package main

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/semustafacevik/go-workspace/go-turkiye/proxy"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to go-workspace | proxy")
	})

	app.Get("/:key/*", ProxyHandler)

	log.Fatal(app.Listen(":3000"))
}
