package routes

import (
	"connectsCall/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func roomRoutes(ctx *fiber.Ctx) error {
	app := ctx.App()

	app.Post("/create", handlers.CreateRoom)

	return nil
}
