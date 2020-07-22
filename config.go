package dexcomClient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	baseUrl = "https://api.dexcom.com"
	devUrl  = "https://sandbox-api.dexcom.com"
	authUrl = "/v1/oauth2/token"
)

type Config struct {
	ClientId     string
	ClientSecret string
	RedirectURI  string
	Sandbox      bool
	IsDev        bool
	IsDebug      bool
	Logging      bool
}

type Token struct {
	AccessToken   string `json:"access_token"`
	ExpiresIn     int64  `json:"expires_in"`
	TokenType     string `json:"token_type"`
	RefreshToken  string `json:"refresh_token"`
	TimeRefreshed int64
}

func (client *DexcomClient) GetOauthToken() (*Token, error) {

	if client.oAuthToken != nil {
		expired := client.oAuthToken.TimeRefreshed + client.oAuthToken.ExpiresIn <= time.Now().Unix()
		if expired {
			goto REQUEST
		}
		return client.oAuthToken, nil
	}

	REQUEST:
	token, err := client.authenticate()
	if err != nil {
		return nil, err
	}
	client.oAuthToken = token
	token.TimeRefreshed = time.Now().Unix()
	return token, err
}

func (client *DexcomClient) GetOAuthTokenWithAuth(authorization string) (*Token, error) {
	client.oAuthToken = &Token{}
	token, err := client.authenticate()
	if err != nil {
		return nil, err
	}
	client.oAuthToken = token
	token.TimeRefreshed = time.Now().Unix()
	return token, err
}

func (client *DexcomClient) SetOAuthToken(token *Token) {
	client.oAuthToken = token
}

func (c *Config) getBaseUrl() string {
	if c.Sandbox {
		return devUrl
	}
	return baseUrl
}

func (client *DexcomClient) authenticate() (*Token, error) {
	req, _ := http.NewRequest("POST", client.config.getBaseUrl()+authUrl, client.getAuthPayload())
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache", "no-cache")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var token Token
	json.Unmarshal(body, &token)

	if token.RefreshToken == "" {
		return nil, errors.New("auth error")
	}

	return &token, nil
}

func (client *DexcomClient) getAuthPayload() *strings.Reader {
	clientSecret := "client_secret=" + client.config.ClientSecret + "&"
	clientId := "client_id=" + client.config.ClientId + "&"
	code := "code=" + client.AuthCode + "&"
	redirectUri := "redirect_uri=" + client.config.RedirectURI
	b := clientSecret + clientId + code + "grant_type=authorization_code&" + redirectUri
	return strings.NewReader(b)
}
