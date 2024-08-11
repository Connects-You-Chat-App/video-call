package server

import (
	"connectsCall/internal/routes"
	"flag"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	port = flag.String("port", os.Getenv("PORT"), "server address")
)

func Run() error {
	flag.Parse()

	fiberApp := fiber.New()

	fiberApp.Use(cors.New())
	fiberApp.Use(logger.New())

	routes.InitializeRoutes(fiberApp)

	return fiberApp.Listen(":" + *port)
}
