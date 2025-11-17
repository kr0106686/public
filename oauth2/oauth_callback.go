package oauth2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/kr0106686/public/domain"
	"github.com/kr0106686/public/oauth2/config"
)

func (o *OAuth) RequestToken(name, code string) (string, error) {
	if name != "kakao" {
		return "", fmt.Errorf("잘못된 요청")
	}
	p := o.cfg

	resp, err := http.PostForm(p.EndPoint.KakaoTokenURL, url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"client_id":     {p.KakaoClientID},
		"client_secret": {p.KakaoClientSecret},
		"redirect_uri":  {p.KakaoRedirectURI},
	})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var t config.Token
	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return "", err
	}
	return t.AccessToken, nil
}

func (o *OAuth) RequestProfile(name, token string) (*domain.User, error) {
	if name != "kakao" {
		return nil, fmt.Errorf("잘못된 요청")
	}
	p := o.cfg

	req, _ := http.NewRequest("GET", p.EndPoint.KakaoInfoURL, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("요청 실패")
	}
	defer resp.Body.Close()

	switch name {
	case "kakao":
		var k config.KakaoProfile
		if err := json.NewDecoder(resp.Body).Decode(&k); err != nil {
			return nil, err
		}
		return &domain.User{
			ProviderID: fmt.Sprintf("%d", k.ID),
			Provider:   "kakao",
			Email:      k.KakaoAccount.Email,
			Name:       k.KakaoAccount.Profile.Nickname,
			Picture:    k.KakaoAccount.Profile.ProfileImageURL,
		}, nil
	}

	return nil, fmt.Errorf("잘못된 요청")
}
