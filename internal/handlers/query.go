package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/alessandrolattao/gosqladmin/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/xwb1989/sqlparser"
)

// QueryHandler returns an echo.HandlerFunc to handle queries for a given database.
func QueryHandler(logger zerolog.Logger, db *database.DB, driverName string) echo.HandlerFunc {
	return func(c echo.Context) error {
		databaseName, query, page, pageSize, err := extractRequestParameters(c, logger)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// Log the data request for debugging purposes
		logger.Debug().
			Str("databaseName", databaseName).
			Str("query", query).
			Int("page", page).
			Int("pageSize", pageSize).
			Msg("Data request")

		isSelect, err := isSelect(query)
		if err != nil {
			logger.Error().Err(err).Msg("Error checking if query is a SELECT statement")
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		logger.Debug().Bool("isSelect", isSelect).Msg("Checked if query is a SELECT statement")

		// Fetch the data based on the query, page, pageSize, and driver
		data, columnInfo, affectedRows, err := db.Query(logger, driverName, query, isSelect, page, pageSize)
		if err != nil {
			logger.Error().Err(err).Msg("Error fetching table data")
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		totalPages := 0
		// Calculate the total number of pages if the query is a SELECT statement
		if isSelect {
			totalPages, err = db.TotalPages(logger, query, pageSize)
			if err != nil {
				logger.Error().Err(err).Msg("Error calculating total pages")
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}

		// Render the data in the template
		return c.Render(http.StatusOK, "data.html", map[string]interface{}{
			"DatabaseName": databaseName,
			"Query":        query,
			"IsSelect":     isSelect,
			"AffectedRows": affectedRows,
			"ColumnInfo":   columnInfo,
			"Data":         data,
			"Page":         page,
			"PageSize":     pageSize,
			"TotalPages":   totalPages,
		})
	}
}

// extractRequestParameters extracts the database name, query, page, and pageSize from the request.
func extractRequestParameters(c echo.Context, logger zerolog.Logger) (string, string, int, int, error) {
	databaseName := c.Param("databasename")
	query := strings.TrimSuffix(getStringFormValue(c, "query"), ";")
	page := getIntFormValue(c, "page", 1)
	pageSize := getIntFormValue(c, "pageSize", 10)

	logger.Debug().
		Str("databaseName", databaseName).
		Str("query", query).
		Int("page", page).
		Int("pageSize", pageSize).
		Msg("Extracted request parameters")

	if databaseName == "" || query == "" {
		return "", "", 0, 0, fmt.Errorf("database name and query must be provided")
	}

	return databaseName, query, page, pageSize, nil
}

// isSelect checks if the given SQL query is a SELECT statement.
func isSelect(query string) (bool, error) {
	// Parse the query using the sqlparser library
	stmt, err := sqlparser.Parse(query)
	if err != nil {
		return false, fmt.Errorf("failed to parse the query: %v", err)
	}

	// Check if the parsed statement is of type *sqlparser.Select
	_, isSelect := stmt.(*sqlparser.Select)
	return isSelect, nil
}

// Helper functions for extracting form values
func getStringFormValue(c echo.Context, key string) string {
	return c.FormValue(key)
}

func getIntFormValue(c echo.Context, key string, defaultValue int) int {
	value := c.FormValue(key)
	if value == "" {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return intValue
}
