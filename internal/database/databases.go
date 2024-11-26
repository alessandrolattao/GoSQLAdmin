package database

import (
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
