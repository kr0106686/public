package filer

import (
	"html/template"
)

const (
	pagePath   = "/templates/page"
	layoutPath = "/templates/layout/*.html"
	docsPath   = "/docs"
	staticPath = "/static"
)

type Filer struct {
	pagePath   string
	layoutTmpl *template.Template
	docsPath   string
	staticPath string
}

func New() *Filer {
	return &Filer{
		pagePath:   pagePath,
		layoutTmpl: template.Must(template.ParseGlob(layoutPath)),
		docsPath:   docsPath,
		staticPath: staticPath,
	}
}
