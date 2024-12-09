package server

import (
	"strconv"

	"github.com/alessandrolattao/gosqladmin/internal/database"
	"github.com/alessandrolattao/gosqladmin/internal/environment"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

// Server represents the Echo server instance.
type Server struct {
	Echo *echo.Echo
}

// NewServer initializes a new Echo server and sets up routes, middleware, and custom logging.
func NewServer(logger zerolog.Logger, db *database.DB, env *environment.Environment) *Server {
	// Create a new Echo instance
	e := echo.New()

	// Manage errors without using JSON
	e.HTTPErrorHandler = CustomHTTPErrorHandler

	// Add RequestLogger middleware to log request details
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Str("method", c.Request().Method).
				Msg("Handled request")
			return nil
		},
	}))

	// Add recovery middleware
	e.Use(middleware.Recover())

	// Set the custom template renderer
	e.Renderer = NewTemplateRenderer(logger)

	// Register routes with driver support
	registerRoutes(e, logger, db, env)

	return &Server{Echo: e}
}

// Start launches the Echo server on the specified port.
// Logs an error if the server fails to start.
func (s *Server) Start(port string) error {
	return s.Echo.Start(":" + port)
}

// getStringFormValue retrieves a string form value or a default if it doesn't exist.
func getStringFormValue(c echo.Context, name string) string {
	return c.FormValue(name)
}

// getIntFormValue retrieves an integer form value or a default if it doesn't exist.
func getIntFormValue(c echo.Context, name string, defaultValue int) int {
	value := c.FormValue(name)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}
