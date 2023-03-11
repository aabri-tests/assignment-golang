package main

import (
	"github.com/aabri-assignments/assignment-golang/handlers"
	"github.com/aabri-assignments/assignment-golang/pkg/logger"
	"github.com/aabri-assignments/assignment-golang/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Logger
	log := logger.InitLogger()

	// Handler
	handler := handlers.NewHandler()

	// Fiber http server
	app := fiber.New()

	// Fiber logger middleware
	app.Use(logger.Middleware(log))

	// Fiber routes
	routes.RegisterRoutes(app, handler)

	// Serve
	log.Error().Err(app.Listen(":8080"))
}
