package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	*zerolog.Logger
}

func InitLogger() *Logger {
	// log time format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// log level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// logger creation
	logger := log.Output(os.Stdout).With().Str("service", "tracker").Logger()

	// global logger setup
	log.Logger = logger

	return &Logger{&logger}
}
