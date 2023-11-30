package config

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLogger() (*os.File, *zerolog.Logger) {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logLevel, debugOn := os.LookupEnv("VAULTY_LOG_LEVEL")

	var logger zerolog.Logger

	// Default level for this example is info, unless debug flag is present
	if debugOn {
		level, err := zerolog.ParseLevel(logLevel)
		if err != nil {
			log.Fatal().Err(err).Msg("Invalid log level")
		}
		zerolog.SetGlobalLevel(level)
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	} else {
		// If debugOn is false, discard all log messages
		logger = zerolog.Nop()
	}

	var logFile *os.File

	// Check if file for logging is set
	LOG_FILE, exists := os.LookupEnv("VAULTY_LOG_FILE")
	if exists {
		logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Panic().Err(err).Msg("Error opening log file")
		}
		logger = logger.Output(zerolog.ConsoleWriter{Out: logFile, TimeFormat: zerolog.TimeFieldFormat})
	}

	return logFile, &logger
}
