package server

import (
	"github.com/alessandrolattao/gosqladmin/internal/database"
	handlers "github.com/alessandrolattao/gosqladmin/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func registerRoutes(e *echo.Echo, logger zerolog.Logger, db *database.DB) {

	// Define routes
	e.GET("/", handlers.HomeHandler)
	e.POST("/dashboard", handlers.DashboardHandler)
	e.POST("/databases", handlers.DatabasesHandler(logger, db))
	e.POST("/tables", handlers.TablesHandler(logger, db))
	e.POST("/table/:databasename/:tablename", handlers.TableHandler())
	e.POST("/query/:databasename", handlers.QueryHandler(logger, db))

	// Serve static files from the "web/static" directory
	e.Static("/static", "web/static")
}
