package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHandler_HttpHealthHandler(t *testing.T) {
	testScenarios := []struct {
		name     string
		method   string
		expected int
	}{
		{
			name:     "GET request returns response 200",
			method:   "GET",
			expected: 200,
		},
		{
			name:     "POST request returns response 405",
			method:   "POST",
			expected: 405,
		},
	}

	path := "/_health"
	handler := NewHandler()
	app := fiber.New()
	app.Get(path, handler.HTTPHealthHandler())

	for _, testScenario := range testScenarios {
		req := httptest.NewRequest(testScenario.method, path, nil)

		resp, _ := app.Test(req, -1) // -1 for no request latency

		assert.Equalf(t, testScenario.expected, resp.StatusCode, testScenario.name)
	}
}
