package routes

import (
	"github.com/aabri-assignments/assignment-golang/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(a *fiber.App, h *handlers.Handler) {
	// Health check endpoint
	a.Get("/_health", h.HTTPHealthHandler())

	// track endpoint
	a.Post("/track", h.TrackHandler())
}
