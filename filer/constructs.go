package filer

import (
	"net/http"

	"github.com/kr0106686/public/oauth2/config"
)

type Interface interface {
	ReadHtmlTmpl(w http.ResponseWriter, name string, data map[string]interface{}) error
	ReadOAuthCfg() (*config.Provider, error)
}
