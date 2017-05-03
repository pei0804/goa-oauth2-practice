package util

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type (
	Github struct {
		AppID  string `toml:"app_id"`
		Secret string `toml:"secret"`
		Scope  string `toml:"scope"`
	}
)

func NewGitHubOAuth(AppID string, Secret string, Scope string) *Github {
	return &Github{
		AppID:  AppID,
		Secret: Secret,
		Scope:  Scope,
	}
}

func (g *Github) GetAuthorizeUrl(state string) string {
	return "https://github.com/login/oauth/authorize" +
		"?client_id=" + g.AppID +
		"&scope=" + url.QueryEscape(g.Scope) +
		"&state=" + url.QueryEscape(state)
}

func (g *Github) GenerateToken(code, state string) (string, error) {
	url := "https://github.com/login/oauth/access_token" +
		"?client_id=" + g.AppID +
		"&client_secret=" + g.Secret +
		"&code=" + code +
		"&state=" + state

	token := map[string]string{}
	err := getJSON(url, &token)
	if _, ok := token["access_token"]; !ok {
		return "", err
	}
	return token["access_token"], nil
}

func (g *Github) GetUserInfo(token string) map[string]interface{} {
	url := "https://api.github.com/user" +
		"?access_token=" + token

	userInfo := map[string]interface{}{}
	getJSON(url, &userInfo)

	// メールアドレスが非公開なら、プライマリメールアドレスを取得
	if _, ok := userInfo["email"]; !ok || userInfo["email"] == nil {
		url := "https://api.github.com/user/emails" +
			"?access_token=" + token

		type email struct {
			Email    string `json:"email"`
			Verified bool   `json:"verified"`
			Primary  bool   `json:"primary"`
		}
		emails := []email{}
		getJSON(url, &emails)

		for i := range emails {
			if emails[i].Primary {
				userInfo["email"] = emails[i].Email
				break
			}
		}
	}

	return userInfo
}

func getJSON(url string, dest interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(dest)
}
