package database

import (
	"fmt"

	"github.com/rs/zerolog"
)

// PaginatedTableData retrieves all columns and their values from a table with pagination
func (db *DB) PaginatedTableData(logger zerolog.Logger, tableName string, page, pageSize int) ([]map[string]interface{}, error) {
	logger.Debug().Msgf("Fetching data from table '%s' with page %d and pageSize %d", tableName, page, pageSize)

	offset := (page - 1) * pageSize
	query := fmt.Sprintf("SELECT * FROM %s LIMIT ? OFFSET ?", tableName)

	rows, err := db.Conn.Queryx(query, pageSize, offset)
	if err != nil {
		logger.Error().Err(err).Msgf("Error executing query for table '%s'", tableName)
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		row := make(map[string]interface{})
		if err := rows.MapScan(row); err != nil {
			logger.Error().Err(err).Msg("Error scanning row into map")
			return nil, err
		}

		// Converti tutti i valori []byte in stringhe
		for key, value := range row {
			if bytes, ok := value.([]byte); ok {
				row[key] = string(bytes)
			}
		}

		results = append(results, row)
	}

	logger.Debug().Msgf("Fetched %d rows from table '%s'", len(results), tableName)
	return results, nil
}

// GetColumnNames retrieves the column names of a specified table
func (db *DB) GetColumnNames(logger zerolog.Logger, tableName string) ([]string, error) {
	logger.Debug().Msgf("Fetching column names for table '%s'", tableName)

	// Query to retrieve column names
	query := fmt.Sprintf("SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = ? AND TABLE_SCHEMA = DATABASE()")

	var columnNames []string
	err := db.Conn.Select(&columnNames, query, tableName)
	if err != nil {
		logger.Error().Err(err).Msgf("Error fetching column names for table '%s'", tableName)
		return nil, err
	}

	logger.Debug().Msgf("Fetched %d column names for table '%s'", len(columnNames), tableName)
	return columnNames, nil
}
