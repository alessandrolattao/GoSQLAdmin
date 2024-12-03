package handlers

import (
	"net/http"

	"github.com/alessandrolattao/gosqladmin/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func DatabasesHandler(logger zerolog.Logger, db *database.DB, driverName string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Fetch list of databases based on the driver
		databaseItems, err := db.ListDatabases(logger, driverName)
		if err != nil {
			logger.Error().Err(err).Msgf("Error fetching list of databases for driver: %s", driverName)
			return echo.NewHTTPError(http.StatusInternalServerError, "Unable to fetch databases")
		}

		// Render the template with data
		return c.Render(http.StatusOK, "databases.html", map[string]interface{}{
			"DatabaseItems": databaseItems,
		})
	}
}
