package dexcomClient

import (
	"io/ioutil"
	"net/http"
	"strings"
	"encoding/json"
)

const BASE_URL = "https://api.dexcom.com"
const DEV_URL = "https://sandbox-api.dexcom.com"
const AUTH_URL = "/v1/oauth2/token"

type Config struct {
	ClientId     string
	ClientSecret string
	IsDev        bool
	Sandbox      bool
	IsDebug      bool
	Logging      bool
	AuthCode     string
	RedirectURI  string
	oAuthToken   *Token
	DexcomToken  string
}

type Token struct {
	AccessToken   string `json:"access_token"`
	ExpiresIn     uint64 `json:"expires_in"`
	TokenType     string `json:"token_type"`
	RefreshToken  string `json:"refresh_token"`
	TimeRefreshed int
}

func (c *Config) GetOauthToken() (*Token, error) {
	if c.oAuthToken != nil {
		return c.oAuthToken, nil
	}

	token, err := c.authenticate()
	if err != nil {
		return nil, err
	}
	c.oAuthToken = token
	return token, err
}

func (c *Config) SetOAuthToken(token *Token) {
	c.oAuthToken = token
}

func (c *Config) GetBaseUrl() string {
	if c.Sandbox {
		return DEV_URL
	}
	return BASE_URL
}

func (c *Config) authenticate() (*Token, error) {
	req, _ := http.NewRequest("POST", c.GetBaseUrl()+AUTH_URL, c.getAuthPayload())
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache", "no-cache")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var token Token
	json.Unmarshal(body, &token)
	return &token, nil
}

func (c *Config) getAuthPayload() *strings.Reader {
	clientSecret := "client_secret=" + c.ClientSecret + "&"
	clientId := "client_id=" + c.ClientId + "&"
	code := "code=" + c.AuthCode + "&"
	redirectUri := "redirect_uri=" + c.RedirectURI
	b := clientSecret + clientId + code + "grant_type=authorization_code&" + redirectUri
	return strings.NewReader(b)
}
