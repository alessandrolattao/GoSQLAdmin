package handlers

import (
	"net/http"

	"github.com/alessandrolattao/gosqladmin/internal/database"
	"github.com/alessandrolattao/gosqladmin/internal/environment"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func DatabasesHandler(logger zerolog.Logger, db *database.DB, env *environment.Environment) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Fetch list of databases based on the driver
		databaseItems, err := db.ListDatabases(logger, env.SQLDriver)
		if err != nil {
			logger.Error().Err(err).Msgf("Error fetching list of databases for driver: %s", env.SQLDriver)
			return echo.NewHTTPError(http.StatusInternalServerError, "Unable to fetch databases")
		}

		// Render the template with data
		return c.Render(http.StatusOK, "databases.html", map[string]interface{}{
			"DatabaseItems":    databaseItems,
			"SelectedDatabase": env.SQLDatabase,
		})
	}
}
