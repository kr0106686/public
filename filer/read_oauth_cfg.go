package filer

import (
	"log"
	"path"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/kr0106686/public/oauth2/config"
)

func (f *Filer) ReadOAuthCfg() (*config.Provider, error) {
	if err := godotenv.Load(path.Join(f.docsPath, "oauth.env")); err != nil {
		log.Print(err)
		return nil, err
	}

	var cfg config.Provider
	if err := env.Parse(&cfg); err != nil {
		log.Print(err)
		return nil, err
	}
	return &cfg, nil
}
