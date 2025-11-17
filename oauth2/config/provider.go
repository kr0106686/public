package config

type (
	Provider struct {
		KakaoClientID     string `env:"KAKAO_CLIENT_ID,required"`
		KakaoClientSecret string `env:"KAKAO_CLIENT_SECRET,required"`
		KakaoRedirectURI  string `env:"KAKAO_REDIRECT_URI,required"`
		EndPoint          EndPoint
	}

	EndPoint struct {
		KakaoAuthURL  string `env:"KAKAO_AUTH_URL,required"`
		KakaoTokenURL string `env:"KAKAO_TOKEN_URL,required"`
		KakaoInfoURL  string `env:"KAKAO_INFO_URL,required"`
	}

	Token struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int64  `json:"expires_in"`
	}

	KakaoProfile struct {
		ID           int64 `json:"id"`
		KakaoAccount struct {
			Email   string `json:"email"`
			Profile struct {
				Nickname        string `json:"nickname"`
				ProfileImageURL string `json:"profile_image_url"`
			} `json:"profile"`
		} `json:"kakao_account"`
	}
)
