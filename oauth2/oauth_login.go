package oauth2

import (
	"fmt"

	"github.com/kr0106686/public/oauth2/config"
)

func (o *OAuth) AuthCodeURL(name string) (string, error) {
	switch name {
	case "kakao":
		return o.createAuthCodeURL(o.cfg), nil
	default:
		return "", fmt.Errorf("지원되지 않는 공급자")
	}
}

func (o *OAuth) createAuthCodeURL(p *config.Provider) string {
	return fmt.Sprintf(
		"%s?client_id=%s&redirect_uri=%s&response_type=code",
		p.EndPoint.KakaoAuthURL,
		p.KakaoClientID,
		p.KakaoRedirectURI,
	)
}
