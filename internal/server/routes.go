package server

import (
	"github.com/alessandrolattao/gosqladmin/internal/database"
	handlers "github.com/alessandrolattao/gosqladmin/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func registerRoutes(e *echo.Echo, logger zerolog.Logger, db *database.DB, driverName string) {
	// Define routes
	e.GET("/", handlers.HomeHandler)
	e.POST("/dashboard", handlers.DashboardHandler)
	e.POST("/databases", handlers.DatabasesHandler(logger, db, driverName))
	e.POST("/tables", handlers.TablesHandler(logger, db, driverName))
	e.POST("/table/:databasename/:tablename", handlers.TableHandler())
	e.POST("/query/:databasename", handlers.QueryHandler(logger, db, driverName))

	// Serve static files from the embedded filesystem
	e.GET("/static/*", StaticFileHandler(logger))
}
