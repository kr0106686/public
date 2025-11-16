package filer

import (
	"fmt"
	"net/http"
	"path"
)

func (f *Filer) ReadHtmlTmpl(w http.ResponseWriter, name string, data map[string]interface{}) error {
	fileName := fmt.Sprintf("%s.html", name)
	filePath := path.Join(f.pagePath, fileName)

	f.layoutTmpl.ParseFiles(filePath)
	return f.layoutTmpl.Execute(w, data)
}
