package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Echo *echo.Echo
}

func NewServer() *Server {
	e := echo.New()

	// Middleware di logging (opzionale)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = NewTemplateRenderer()

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

	e.POST("/tables", func(c echo.Context) error {
		// Array di stringhe
		menuItems := []string{"Home", "About Us", "Services", "Contact"}

		// Renderizza il template con i dati
		return c.Render(http.StatusOK, "tables.html", map[string]interface{}{
			"MenuItems": menuItems,
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

	e.Static("/static", "web/static")

	return &Server{Echo: e}
}

func (s *Server) Start(port string) error {
	return s.Echo.Start(":" + port)
}
