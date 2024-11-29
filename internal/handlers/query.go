package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/alessandrolattao/gomyadmin/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

// QueryHandler returns an echo.HandlerFunc to handle queries for a given database.
func QueryHandler(logger zerolog.Logger, db *database.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract the dynamic part of the URL (database name)
		databaseName := c.Param("databasename")

		// Extract form values for query, page, and pageSize with defaults
		query := getStringFormValue(c, "query")
		page := getIntFormValue(c, "page", 1)
		pageSize := getIntFormValue(c, "pageSize", 10)

		// Remove any trailing semicolons from the query
		query = strings.TrimSuffix(query, ";")

		// Log the data request for debugging purposes
		logger.Debug().
			Str("databaseName", databaseName).
			Str("query", query).
			Int("page", page).
			Int("pageSize", pageSize).
			Msg("Data request")

		// Select the database for the operation
		db.SelectDatabase(logger, databaseName)

		// Fetch the total number of pages for the query
		totalPages, err := db.TotalPages(logger, query, pageSize)
		if err != nil {
			logger.Error().Err(err).Msg("Error fetching total pages")
			return err
		}

		// Fetch column names for the query
		columnNames, err := db.GetColumnNames(logger, query)
		if err != nil {
			logger.Error().Err(err).Msg("Error fetching column names")
			return err
		}

		// Fetch the data based on the query, page, and pageSize
		data, err := db.Query(logger, query, page, pageSize)
		if err != nil {
			logger.Error().Err(err).Msg("Error fetching table data")
			return err
		}

		// Render the data in the template
		return c.Render(http.StatusOK, "data.html", map[string]interface{}{
			"DatabaseName": databaseName,
			"Query":        query,
			"ColumnNames":  columnNames,
			"Data":         data,
			"Page":         page,
			"PageSize":     pageSize,
			"TotalPages":   totalPages,
		})
	}
}

// Helper functions for extracting form values (assumed to be defined elsewhere)
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
