package web

import "embed"

//go:embed static/*
var StaticFiles embed.FS

//go:embed templates/*
var TemplateFiles embed.FS
