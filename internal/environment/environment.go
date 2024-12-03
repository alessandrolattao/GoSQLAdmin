package environment

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
)

type Environment struct {
	SQLUser         string
	SQLPassword     string
	SQLHost         string
	SQLPort         string
	SQLConnTimeout  time.Duration
	SQLReadTimeout  time.Duration
	SQLWriteTimeout time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

// getEnvOrError fetches an environment variable and returns an error if it is not set
func getEnvOrError(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %s is not set", key)
	}
	return value, nil
}

// getEnvOrDefault fetches an environment variable and returns a default value if it is not set
func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// parseEnvOrDefault converts an environment variable to an integer or returns a default value
func parseEnvOrDefault(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return parsedValue
}

// parseDurationEnvOrDefault converts an environment variable to a time.Duration or returns a default value
func parseDurationEnvOrDefault(key string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	parsedValue, err := time.ParseDuration(value)
	if err != nil {
		return defaultValue
	}
	return parsedValue
}

// GetEnvironment initializes and validates the environment configuration
func GetEnvironment(logger zerolog.Logger) (*Environment, error) {
	env := &Environment{}

	var err error
	// Fetch and validate required environment variables
	if env.SQLUser, err = getEnvOrError("SQL_USER"); err != nil {
		return nil, err
	}
	if env.SQLPassword, err = getEnvOrError("SQL_PASSWORD"); err != nil {
		return nil, err
	}
	if env.SQLHost, err = getEnvOrError("SQL_HOST"); err != nil {
		return nil, err
	}

	// Fetch and validate optional environment variables with defaults
	env.SQLPort = getEnvOrDefault("SQL_PORT", "3306")
	env.SQLConnTimeout = parseDurationEnvOrDefault("SQL_CONN_TIMEOUT", 60*time.Second)
	env.SQLReadTimeout = parseDurationEnvOrDefault("SQL_READ_TIMEOUT", 30*time.Minute)
	env.SQLWriteTimeout = parseDurationEnvOrDefault("SQL_WRITE_TIMEOUT", 30*time.Minute)

	// Fetch and validate connection pool settings
	env.MaxOpenConns = parseEnvOrDefault("SQL_MAX_OPEN_CONNS", 5)
	env.MaxIdleConns = parseEnvOrDefault("SQL_MAX_IDLE_CONNS", 5)
	env.ConnMaxLifetime = parseDurationEnvOrDefault("SQL_CONN_MAX_LIFETIME", 30*time.Minute)

	logger.Debug().
		Str("sql_user", env.SQLUser).
		Str("sql_host", env.SQLHost).
		Str("sql_port", env.SQLPort).
		Int("max_open_conns", env.MaxOpenConns).
		Int("max_idle_conns", env.MaxIdleConns).
		Dur("conn_max_lifetime", env.ConnMaxLifetime).
		Msg("Environment variables successfully loaded")

	return env, nil
}
