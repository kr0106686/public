package filer

const (
	pagePath   = "/templates/page"
	layoutPath = "/templates/layout/*.html"
	docsPath   = "/docs"
	staticPath = "/static"
)

type Filer struct {
	pagePath   string
	layoutPath string
}

func New() *Filer {
	return &Filer{
		pagePath:   pagePath,
		layoutPath: layoutPath,
	}
}
