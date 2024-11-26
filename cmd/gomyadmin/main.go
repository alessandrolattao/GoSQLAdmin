package main

import (
	"os"
	"strings"

	"github.com/alessandrolattao/gomyadmin/internal/database"
	"github.com/alessandrolattao/gomyadmin/internal/environment"
	"github.com/alessandrolattao/gomyadmin/internal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// parseLogLevel parses the log level from an environment variable or defaults to INFO.
func parseLogLevel(envVar string) zerolog.Level {
	levelStr := strings.ToLower(envVar)
	switch levelStr {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel // Default log level
	}
}

func main() {

	// Set log level based on LOG_LEVEL environment variable
	logLevel := parseLogLevel(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(logLevel)

	// Get the server port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Configure zerolog for console output
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Get environment variables for database connection
	env, err := environment.GetEnvironment(log.Logger)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get environment variables")
	}

	db, err := database.NewDB(log.Logger, env)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database")
	}
	defer db.Close()

	// Initialize the server with the logger and database connection
	srv := server.NewServer(log.Logger, db)
	log.Info().Msgf("Starting GoMyAdmin on http://localhost:%s", port)

	// Start the server and log critical errors if any
	if err := srv.Start(port); err != nil {
		log.Fatal().Err(err).Msg("Error starting server")
	}
}
