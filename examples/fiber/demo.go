package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to go-workspace!")
	})

	app.Get("/w/:name", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome " + ctx.Params("name"))
	})

	app.Get("/opt/:name?", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome " + ctx.Params("name", "GOPHER :)"))
	})

	app.Get("/api/*", getApiPath)

	app.Listen(":3000")
}

func getApiPath(ctx *fiber.Ctx) error {
	return ctx.SendString("API path: " + ctx.Params("*"))
}
