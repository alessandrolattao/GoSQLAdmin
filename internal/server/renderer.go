package server

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

// TemplateRenderer is a custom renderer for Echo, based on Go's html/template package.
type TemplateRenderer struct {
	templates *template.Template
	logger    zerolog.Logger
}

// NewTemplateRenderer initializes a new TemplateRenderer with the given logger.
// It loads HTML templates from the specified directory and adds custom functions.
func NewTemplateRenderer(logger zerolog.Logger) *TemplateRenderer {
	funcMap := template.FuncMap{
		"add":      func(a, b int) int { return a + b },
		"subtract": func(a, b int) int { return a - b },
	}

	templates, err := template.New("").Funcs(funcMap).ParseGlob("web/templates/*.html")
	if err != nil {
		logger.Fatal().Err(err).Msg("Error loading templates")
	}
	return &TemplateRenderer{templates: templates, logger: logger}
}

// Render renders a template with the given name and data.
// Logs the rendering process and any errors encountered.
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := t.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		t.logger.Error().Err(err).Str("template", name).Msg("Error rendering template")
	}
	return err
}
