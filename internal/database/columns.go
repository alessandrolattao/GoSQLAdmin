package database

import (
	"fmt"

	"github.com/rs/zerolog"
)

// GetResultColumnNames retrieves the column names from the result of a query
func (db *DB) GetColumnNames(logger zerolog.Logger, query string) ([]string, error) {
	logger.Debug().Msgf("Fetching result column names for query: '%s'", query)

	// Add a LIMIT 0 to the query to fetch only metadata
	limitedQuery := fmt.Sprintf("SELECT * FROM (%s) AS temp LIMIT 0", query)

	// Execute the query and fetch metadata
	rows, err := db.Conn.Queryx(limitedQuery)
	if err != nil {
		logger.Error().Err(err).Msg("Error executing query for column metadata")
		return nil, err
	}
	defer rows.Close()

	// Retrieve column names from metadata
	columnNames, err := rows.Columns()
	if err != nil {
		logger.Error().Err(err).Msg("Error retrieving column names from metadata")
		return nil, err
	}

	logger.Debug().Msgf("Fetched %d column names for query: '%s'", len(columnNames), query)
	return columnNames, nil
}
