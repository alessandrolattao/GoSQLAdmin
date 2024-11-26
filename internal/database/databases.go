package database

import (
	"fmt"

	"github.com/rs/zerolog"
)

// ListDatabases retrieves and returns a list of all databases
func (db *DB) ListDatabases(logger zerolog.Logger) ([]string, error) {
	logger.Debug().Msg("Fetching list of databases")

	// Query to list all databases
	query := "SHOW DATABASES"

	var databases []string
	err := db.Conn.Select(&databases, query)
	if err != nil {
		logger.Error().Err(err).Msg("Error fetching list of databases")
		return nil, err
	}

	logger.Debug().Msgf("Fetched %d databases", len(databases))
	return databases, nil
}

// SelectDatabase sets the active database for the connection
func (db *DB) SelectDatabase(logger zerolog.Logger, databaseName string) error {
	logger.Debug().Msgf("Selecting database '%s'", databaseName)

	// Query to use the specified database
	query := fmt.Sprintf("USE %s", databaseName)

	_, err := db.Conn.Exec(query)
	if err != nil {
		logger.Error().Err(err).Msgf("Error selecting database '%s'", databaseName)
		return err
	}

	logger.Debug().Msgf("Successfully selected database '%s'", databaseName)
	return nil
}
