package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func DashboardHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "dashboard.html", map[string]interface{}{
		"Title": "Dashboard",
	})

}
