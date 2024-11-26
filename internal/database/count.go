package database

import (
	"fmt"
	"math"

	"github.com/rs/zerolog"
)

// TotalPages calculates and returns the total number of pages for a table based on the given page size
func (db *DB) TotalPages(logger zerolog.Logger, databaseName string, tableName string, pageSize int) (int, error) {
	logger.Debug().Msgf("Calculating total pages for table: %s in database: %s with page size: %d", tableName, databaseName, pageSize)

	// Query to count rows in the specified table
	query := fmt.Sprintf("SELECT COUNT(*) FROM `%s`.`%s`", databaseName, tableName)

	var rowCount int64
	err := db.Conn.Get(&rowCount, query)
	if err != nil {
		logger.Error().Err(err).Msgf("Error counting rows in table: %s from database: %s", tableName, databaseName)
		return 0, err
	}

	// Calculate the total number of pages
	totalPages := int(math.Ceil(float64(rowCount) / float64(pageSize)))

	logger.Debug().Msgf("Table: %s has %d rows, resulting in %d pages with page size: %d", tableName, rowCount, totalPages, pageSize)
	return totalPages, nil
}
