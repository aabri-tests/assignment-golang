package handlers

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHandler_TrackHandler(t *testing.T) {
	testScenarios := []struct {
		name           string
		requestBody    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Successful response",
			requestBody:    `[{"source": "IND","destination": "EWR"}]`,
			expectedStatus: 200,
			expectedBody:   `{"source":"IND","destination":"EWR"}`,
		},
		{
			name:           "Invalid payload",
			requestBody:    `{"foo": "bar"}`,
			expectedStatus: 400,
			expectedBody:   `json: cannot unmarshal object into Go value of type []tracker.Flight`,
		},
	}

	path := "/track"
	handler := NewHandler()
	app := fiber.New()
	app.Post(path, handler.TrackHandler())

	for _, testScenario := range testScenarios {
		req := httptest.NewRequest("POST", path, strings.NewReader(testScenario.requestBody))
		resp, _ := app.Test(req, -1) // -1 for no request latency

		body, _ := io.ReadAll(resp.Body)

		assert.Equalf(t, testScenario.expectedStatus, resp.StatusCode, testScenario.name)
		assert.Equalf(t, testScenario.expectedBody, string(body), testScenario.name)
	}
}
