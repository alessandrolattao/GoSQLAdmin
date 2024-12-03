package database

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog"
)

// ListDatabases retrieves and returns a list of all databases.
// It supports multiple database types (MySQL, PostgreSQL, SQLite, etc.).
func (db *DB) ListDatabases(logger zerolog.Logger, driverName string) ([]string, error) {
	logger.Debug().Msgf("Fetching list of databases using driver: %s", driverName)

	var query string

	// Construct the query based on the database driver
	switch strings.ToLower(driverName) {
	case "mysql":
		query = "SHOW DATABASES"
	case "postgres":
		query = "SELECT datname FROM pg_database WHERE datistemplate = false"
	case "sqlite":
		logger.Warn().Msg("SQLite does not support listing multiple databases")
		return []string{"default"}, nil // Return a default database
	case "sqlserver":
		query = "SELECT name FROM sys.databases"
	case "snowflake":
		query = "SHOW DATABASES"
	case "clickhouse":
		query = "SELECT name FROM system.databases"
	default:
		logger.Warn().Msgf("Unsupported driver: %s", driverName)
		return []string{}, nil // Return an empty list
	}

	var databases []string
	err := db.Conn.Select(&databases, query)
	if err != nil {
		logger.Error().Err(err).Msg("Error fetching list of databases")
		return nil, err
	}

	logger.Debug().Msgf("Fetched %d databases", len(databases))
	return databases, nil
}

// SelectDatabase sets the active database for the connection.
// It supports multiple database types (MySQL, PostgreSQL, etc.).
func (db *DB) SelectDatabase(logger zerolog.Logger, driverName, databaseName string) {
	logger.Debug().Msgf("Selecting database '%s' using driver: %s", databaseName, driverName)

	if strings.TrimSpace(databaseName) == "" {
		logger.Warn().Msg("Database name is empty; no action taken")
		return
	}

	var query string

	// Construct the query based on the database driver
	switch strings.ToLower(driverName) {
	case "mysql":
		query = fmt.Sprintf("USE `%s`", databaseName)
	case "postgres":
		query = fmt.Sprintf("SET search_path TO %s", databaseName)
	case "sqlserver":
		query = fmt.Sprintf("USE [%s]", databaseName)
	case "snowflake":
		query = fmt.Sprintf("USE DATABASE %s", databaseName)
	case "sqlite", "clickhouse":
		logger.Warn().Msgf("Driver %s does not support switching databases dynamically", driverName)
		return
	default:
		logger.Warn().Msgf("Unsupported driver: %s; no action taken", driverName)
		return
	}

	_, err := db.Conn.Exec(query)
	if err != nil {
		logger.Error().Err(err).Msgf("Error selecting database '%s'", databaseName)
		return
	}

	logger.Debug().Msgf("Successfully selected database '%s'", databaseName)
}
