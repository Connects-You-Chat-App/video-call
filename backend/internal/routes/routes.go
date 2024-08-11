package routes

import "github.com/gofiber/fiber/v2"

func InitializeRoutes(app *fiber.App) {
	app.Use("/room", roomRoutes)

}
