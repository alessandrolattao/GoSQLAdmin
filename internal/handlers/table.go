package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// TableHandler returns an echo.HandlerFunc for rendering table details with database and table names.
func TableHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the dynamic parts of the URL
		databaseName := c.Param("databasename")
		tableName := c.Param("tablename")

		// Render the template with the database and table names
		return c.Render(http.StatusOK, "table.html", map[string]interface{}{
			"Title":        "Table",
			"DatabaseName": databaseName,
			"TableName":    tableName,
		})
	}
}
