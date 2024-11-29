package database

import (
	"fmt"

	"github.com/rs/zerolog"
)

// Query execute queries
func (db *DB) Query(logger zerolog.Logger, query string, page, pageSize int) ([]map[string]interface{}, error) {
	offset := (page - 1) * pageSize
	queryString := fmt.Sprintf("%s LIMIT ? OFFSET ?", query)
	logger.Debug().Msgf("Executing query '%s'", queryString)

	rows, err := db.Conn.Queryx(queryString, pageSize, offset)
	if err != nil {
		logger.Error().Err(err).Msgf("Error executing query '%s'", queryString)
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

	logger.Debug().Msgf("Query %s fetched %d rows", queryString, len(results))
	return results, nil
}
