package database

import (
	"github.com/rs/zerolog"
)

type ColumnInfo struct {
	Name string
	Type string
}

// Query executes a query and delegates to appropriate handler based on query type
func (db *DB) Query(logger zerolog.Logger, query string, isSelect bool, page, pageSize int) ([]map[string]interface{}, []ColumnInfo, int, error) {
	if isSelect {
		return db.executeSelectQuery(logger, query, page, pageSize)
	}
	return db.executeNonSelectQuery(logger, query)
}

// executeSelectQuery executes a SELECT query and returns the results
func (db *DB) executeSelectQuery(logger zerolog.Logger, query string, page, pageSize int) ([]map[string]interface{}, []ColumnInfo, int, error) {
	offset := (page - 1) * pageSize
	queryString := query + " LIMIT ? OFFSET ?"
	logger.Debug().Msgf("Executing SELECT query '%s'", queryString)

	// Execute the query
	rows, err := db.Conn.Queryx(queryString, pageSize, offset)
	if err != nil {
		logger.Error().Err(err).Msgf("Error executing SELECT query '%s'", queryString)
		return nil, nil, 0, err
	}
	defer rows.Close()

	// Fetch column names and types
	columnNames, err := rows.Columns()
	if err != nil {
		logger.Error().Err(err).Msg("Error fetching column names for SELECT query")
		return nil, nil, 0, err
	}

	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		logger.Error().Err(err).Msg("Error fetching column types for SELECT query")
		return nil, nil, 0, err
	}

	var columnInfo []ColumnInfo
	for i, col := range columnNames {
		dataType := "unknown"
		if columnTypes[i] != nil {
			dataType = columnTypes[i].DatabaseTypeName()
		}
		columnInfo = append(columnInfo, ColumnInfo{
			Name: col,
			Type: dataType,
		})
	}

	// Fetch rows and map them to a slice of maps
	var results []map[string]interface{}
	for rows.Next() {
		row := make(map[string]interface{})
		if err := rows.MapScan(row); err != nil {
			logger.Error().Err(err).Msg("Error scanning row for SELECT query")
			return nil, nil, 0, err
		}

		// Convert []byte to string for compatibility
		for key, value := range row {
			if bytes, ok := value.([]byte); ok {
				row[key] = string(bytes)
			}
		}

		results = append(results, row)
	}

	if err := rows.Err(); err != nil {
		logger.Error().Err(err).Msg("Error iterating rows for SELECT query")
		return nil, nil, 0, err
	}

	logger.Debug().Msgf("SELECT query '%s' fetched %d rows", queryString, len(results))
	return results, columnInfo, len(results), nil
}

// executeNonSelectQuery executes a non-SELECT query and returns the number of affected rows
func (db *DB) executeNonSelectQuery(logger zerolog.Logger, query string) ([]map[string]interface{}, []ColumnInfo, int, error) {
	logger.Debug().Msgf("Executing non-SELECT query '%s'", query)

	// Execute the query and fetch the number of affected rows
	result, err := db.Conn.Exec(query)
	if err != nil {
		logger.Error().Err(err).Msgf("Error executing non-SELECT query '%s'", query)
		return nil, nil, 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error().Err(err).Msg("Error fetching rows affected by non-SELECT query")
		return nil, nil, 0, err
	}

	logger.Debug().Msgf("Non-SELECT query '%s' affected %d rows", query, rowsAffected)
	return nil, nil, int(rowsAffected), nil
}
