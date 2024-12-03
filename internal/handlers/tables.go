package handlers

import (
	"net/http"

	"github.com/alessandrolattao/gosqladmin/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

// TablesHandler returns an echo.HandlerFunc with injected logger and database dependencies.
func TablesHandler(logger zerolog.Logger, db *database.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract the selected database from the form
		selectedDatabase := c.FormValue("selectedDatabase")
		if selectedDatabase == "" {
			logger.Warn().Msg("No database selected")
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "No database selected",
			})
		}

		// Fetch tables for the selected database
		tableItems, err := db.ListTables(logger, selectedDatabase)
		if err != nil {
			logger.Error().Err(err).Msg("Error fetching list of tables")
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Unable to fetch tables for the selected database",
			})
		}

		// Log successful retrieval of tables
		logger.Debug().
			Str("selectedDatabase", selectedDatabase).
			Int("tableCount", len(tableItems)).
			Msg("Fetched tables successfully")

		// Render the template with the retrieved data
		return c.Render(http.StatusOK, "tables.html", map[string]interface{}{
			"SelectedDatabase": selectedDatabase,
			"TableItems":       tableItems,
		})
	}
}
