package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CustomHTTPErrorHandler is a custom error handler for Echo.
func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var message string

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if he.Message != nil {
			message = he.Message.(string)
		} else {
			message = http.StatusText(code)
		}
	} else {
		message = http.StatusText(code)
	}

	// Send a simple text response
	if !c.Response().Committed {
		c.String(code, message)
	}
}
