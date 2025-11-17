package oauth2

import (
	"github.com/kr0106686/public/domain"
	"github.com/kr0106686/public/filer"
	"github.com/kr0106686/public/oauth2/config"
)

type (
	Interface interface {
		AuthURL(name string) (string, error)
		RequestToken(name, code string) (string, error)
		RequestProfile(name, token string) (*domain.User, error)
	}

	OAuth struct {
		cfg *config.Provider
	}
)

func NewOAuth2() *OAuth {
	f := filer.New()
	cfg, _ := f.ReadOAuthCfg()
	return &OAuth{cfg: cfg}
}
