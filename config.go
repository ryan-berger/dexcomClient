package dexcomClient

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

func (c *Config) getBaseUrl() string {
	if c.Sandbox {
		return devUrl
	}
	return baseUrl
}