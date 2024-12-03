package handlers

import (
	"net/http"

	"github.com/alessandrolattao/gosqladmin/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func DatabasesHandler(logger zerolog.Logger, db *database.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Array of database items
		databaseItems, err := db.ListDatabases(logger)
		if err != nil {
			logger.Error().Err(err).Msg("Error fetching list of databases")
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// Render the template with data
		return c.Render(http.StatusOK, "databases.html", map[string]interface{}{
			"DatabaseItems": databaseItems,
		})
	}
}
