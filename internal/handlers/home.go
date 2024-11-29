package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "base.html", map[string]interface{}{})
}
