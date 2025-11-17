package filer

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func (f *Filer) ReadHtmlTmpl(w http.ResponseWriter, name string, data map[string]interface{}) error {

	fileName := fmt.Sprintf("%s.html", name)
	filePath := path.Join(f.pagePath, fileName)

	return template.Must(f.readLayoutTmpl().ParseFiles(filePath)).Execute(w, data)
}

func (f *Filer) readLayoutTmpl() *template.Template {
	return template.Must(template.ParseGlob(f.layoutPath))
}
