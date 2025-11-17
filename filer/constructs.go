package filer

import (
	"github.com/kr0106686/public/oauth2/config"
)

type Interface interface {
	ReadOAuthCfg() (*config.Provider, error)
	ReadFreeYaml(fileName string) map[string]interface{}
}
