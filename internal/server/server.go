package server

import (
	"net/http"

	"github.com/alessandrolattao/gomyadmin/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

// Server represents the Echo server instance.
type Server struct {
	Echo *echo.Echo
}

// NewServer initializes a new Echo server and sets up routes, middleware, and custom logging.
func NewServer(logger zerolog.Logger, db *database.DB) *Server {

	// Create a new Echo instance
	e := echo.New()

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

	// Define routes
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "base.html", map[string]interface{}{
			"Title": "GoMyAdmin",
		})
	})

	e.POST("/dashboard", func(c echo.Context) error {
		return c.Render(http.StatusOK, "dashboard.html", map[string]interface{}{
			"Title": "Dashboard",
		})
	})

	e.POST("/databases", func(c echo.Context) error {
		// Array of database items
		databaseItems, err := db.ListDatabases(logger)
		if err != nil {
			logger.Error().Err(err).Msg("Error fetching list of databases")
			return err
		}

		// Render the template with data
		return c.Render(http.StatusOK, "databases.html", map[string]interface{}{
			"DatabaseItems": databaseItems,
		})
	})

	e.POST("/tables", func(c echo.Context) error {
		selectedDatabase := c.FormValue("selectedDatabase")

		tableItems, err := db.ListTables(logger, selectedDatabase)
		if err != nil {
			logger.Error().Err(err).Msg("Error fetching list of tables")
			return err
		}

		// Render the template with data
		return c.Render(http.StatusOK, "tables.html", map[string]interface{}{
			"TableItems": tableItems,
		})
	})

	e.POST("/table/:tablename", func(c echo.Context) error {
		// Get the dynamic part of the URL
		tableName := c.Param("tablename")

		return c.Render(http.StatusOK, "table.html", map[string]interface{}{
			"Title":     "Table",
			"TableName": tableName,
		})
	})

	// Serve static files from the "web/static" directory
	e.Static("/static", "web/static")

	return &Server{Echo: e}
}

// Start launches the Echo server on the specified port.
// Logs an error if the server fails to start.
func (s *Server) Start(port string) error {
	return s.Echo.Start(":" + port)
}
