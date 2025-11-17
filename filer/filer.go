package filer

const (
	pagePath   = "/templates/page"
	layoutPath = "/templates/layout/*.html"
	docsPath   = "/docs"
)

type Filer struct {
	pagePath   string
	layoutPath string
	docsPath   string
}

func New() Interface {
	return &Filer{
		docsPath:   docsPath,
		pagePath:   pagePath,
		layoutPath: layoutPath,
	}
}
