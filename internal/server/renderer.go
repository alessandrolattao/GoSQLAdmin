package server

import (
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer() *TemplateRenderer {
	templates, err := template.ParseGlob("web/templates/*.html")
	if err != nil {
		log.Fatalf("Error loading templates: %v", err)
	}

	return &TemplateRenderer{templates: templates}
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
