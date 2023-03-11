package logger

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const (
	LocalsKey             = "logger-locals-key"
	CorrelationHeaderName = fiber.HeaderXRequestID
)

func Middleware(logger *Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		l := logger.With().Str("correlation-id", extractCorrelationID(logger, c)).Logger()

		c.Locals(LocalsKey, &l)

		start := time.Now()

		msg := ""

		err := c.Next()
		if err != nil {
			msg = err.Error()
			_ = c.SendStatus(fiber.StatusInternalServerError)
		}

		ll := l.With().
			Int("status", c.Response().StatusCode()).
			Str("protocol", c.Protocol()).
			Str("method", c.Method()).
			Str("host", c.Hostname()).
			Str("path", c.Path()).
			Str("ip", c.IP()).
			Str("latency", time.Since(start).String()).
			Str("user-agent", c.Get(fiber.HeaderUserAgent)).
			Logger()

		switch {
		case c.Response().StatusCode() >= fiber.StatusBadRequest && c.Response().StatusCode() < fiber.StatusInternalServerError:
			ll.Warn().Msg(msg)
		case c.Response().StatusCode() >= http.StatusInternalServerError:
			ll.Error().Msg(msg)
		default:
			ll.Info().Msg(msg)
		}

		return err
	}
}

func extractCorrelationID(logger *Logger, c *fiber.Ctx) string {
	if corrID := c.Get(CorrelationHeaderName, ""); corrID != "" {
		return corrID
	}

	corrID, err := uuid.NewRandom()

	if err != nil {
		logger.Error().Stack().Err(err).Msg("error generating request id")
	}

	return corrID.String()
}
