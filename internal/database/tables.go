package database

import (
	"fmt"

	"github.com/rs/zerolog"
)

// ListTables retrieves and returns a list of all tables in the specified database
func (db *DB) ListTables(logger zerolog.Logger, databaseName string) ([]string, error) {
	logger.Debug().Msgf("Fetching list of tables for database: %s", databaseName)

	// Query to list all tables in the specified database
	query := fmt.Sprintf("SHOW TABLES FROM `%s`", databaseName)

	var tables []string
	err := db.Conn.Select(&tables, query)
	if err != nil {
		logger.Error().Err(err).Msgf("Error fetching list of tables for database: %s", databaseName)
		return nil, err
	}

	logger.Debug().Msgf("Fetched %d tables from database: %s", len(tables), databaseName)
	return tables, nil
}
