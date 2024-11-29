package database

import (
	"fmt"
	"math"

	"github.com/rs/zerolog"
)

// TotalPages calculates and returns the total number of pages for a query based on the given page size
func (db *DB) TotalPages(logger zerolog.Logger, query string, pageSize int) (int, error) {

	logger.Debug().Msgf("Calculating total pages for query: %s with page size: %d", query, pageSize)

	// Wrap the original query to count total rows
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM (%s) AS subquery", query)

	var rowCount int64
	if err := db.Conn.Get(&rowCount, countQuery); err != nil {
		logger.Error().Err(err).Msgf("Error counting rows for query: %s", query)
		return 0, err
	}

	// Calculate the total number of pages
	totalPages := int(math.Ceil(float64(rowCount) / float64(pageSize)))

	logger.Debug().Msgf("Query has %d rows, resulting in %d pages with page size: %d", rowCount, totalPages, pageSize)

	return totalPages, nil

}
