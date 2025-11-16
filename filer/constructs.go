package filer

import "net/http"

type Interface interface {
	ReadHtmlTmpl(w http.ResponseWriter, name string, data map[string]interface{}) error
}
