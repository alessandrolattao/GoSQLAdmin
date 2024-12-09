package server

import (
	"github.com/alessandrolattao/gosqladmin/internal/database"
	"github.com/alessandrolattao/gosqladmin/internal/environment"
	handlers "github.com/alessandrolattao/gosqladmin/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func registerRoutes(e *echo.Echo, logger zerolog.Logger, db *database.DB, env *environment.Environment) {
	// Define routes
	e.GET("/", handlers.HomeHandler)
	e.POST("/dashboard", handlers.DashboardHandler)
	e.POST("/databases", handlers.DatabasesHandler(logger, db, env))
	e.POST("/tables", handlers.TablesHandler(logger, db, env.SQLDriver))
	e.POST("/table/:databasename/:tablename", handlers.TableHandler())
	e.POST("/query/:databasename", handlers.QueryHandler(logger, db, env.SQLDriver))

	// Serve static files from the embedded filesystem
	e.GET("/static/*", StaticFileHandler(logger))
}
