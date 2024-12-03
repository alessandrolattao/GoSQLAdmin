package database

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog"
)

// ListTables retrieves and returns a list of all tables in the specified database.
// It supports multiple database types (MySQL, PostgreSQL, SQLite, etc.).
func (db *DB) ListTables(logger zerolog.Logger, driverName, databaseName string) ([]string, error) {
	logger.Debug().Msgf("Fetching list of tables for database: %s using driver: %s", databaseName, driverName)

	// Validate input
	if strings.TrimSpace(databaseName) == "" && driverName != "sqlite" {
		err := fmt.Errorf("database name cannot be empty for driver: %s", driverName)
		logger.Error().Err(err).Msg("Error in ListTables")
		return nil, err
	}

	var query string

	// Construct the query based on the database driver
	switch strings.ToLower(driverName) {
	case "mysql":
		query = fmt.Sprintf("SHOW TABLES FROM `%s`", databaseName)
	case "postgres":
		query = "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_catalog = $1"
	case "sqlite":
		query = "SELECT name FROM sqlite_master WHERE type='table'"
	case "sqlserver":
		query = "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_TYPE = 'BASE TABLE'"
	case "snowflake":
		query = fmt.Sprintf("SHOW TABLES IN SCHEMA %s", databaseName)
	case "clickhouse":
		query = "SELECT name FROM system.tables WHERE database = ?"
	default:
		err := fmt.Errorf("unsupported driver: %s", driverName)
		logger.Error().Err(err).Msg("Error in ListTables")
		return nil, err
	}

	var tables []string
	var err error

	// Execute the query with database-specific parameters if needed
	switch strings.ToLower(driverName) {
	case "postgres", "clickhouse":
		err = db.Conn.Select(&tables, query, databaseName)
	default:
		err = db.Conn.Select(&tables, query)
	}

	if err != nil {
		logger.Error().Err(err).Msgf("Error fetching list of tables for database: %s", databaseName)
		return nil, err
	}

	logger.Debug().Msgf("Fetched %d tables from database: %s", len(tables), databaseName)
	return tables, nil
}
