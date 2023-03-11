package handlers

import (
	"bytes"

	"github.com/aabri-assignments/assignment-golang/internal/tracker"
	"github.com/gofiber/fiber/v2"
)

var flights []tracker.Flight

type ErrorResponse struct {
	Message string `json:"message"`
}

func (h *Handler) TrackHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqBody := c.Body()

		errChan := make(chan error)
		t := make(chan tracker.Flight)

		// Start a new goroutine to handle the request
		go h.handleRequest(reqBody, t, errChan)

		select {
		case err := <-errChan:
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		case flight := <-t:
			return c.JSON(flight)
		}
	}
}

func (h *Handler) handleRequest(reqBody []byte, t chan tracker.Flight, errChan chan error) {
	err := h.dec.Decode(bytes.NewReader(reqBody), &flights)
	if err != nil {
		errChan <- err

		return
	}

	newTracker := tracker.New(flights)

	t <- newTracker.Track()
}
