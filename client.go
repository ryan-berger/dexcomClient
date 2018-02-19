package dexcomClient

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type DexcomClient struct {
	config *Config
	*AuthClient
	*EstimatedGlucoseClient
}

func NewClient(config *Config) *DexcomClient {
	dc := &DexcomClient{
		config:                 config,
		AuthClient:             NewAuthClient(config),
		EstimatedGlucoseClient: NewEGVClient(config, &defaultLogger{config: config}),
	}

	if config.IsDev {
		fmt.Println("Dev server starting on :8000")
		fmt.Println(config.getBaseUrl() + "/v1/oauth2/login?client_id=" + config.ClientId + "&redirect_uri=" + config.RedirectURI + "&response_type=code&scope=offline_access")
		defer dc.startDevServer()
	}
	return dc
}

func NewClientWithToken(client *http.Client, config *Config, token *Token) *DexcomClient {
	config.SetOAuthToken(token)
	return &DexcomClient{
		config:                 config,
		AuthClient:             NewAuthClient(config),
		EstimatedGlucoseClient: NewEGVClient(config, &defaultLogger{config: config}),
	}
}

func (client *DexcomClient) startDevServer() {
	server := &http.Server{Addr: ":8000"}

	router := mux.NewRouter()
	router.Path("/oauth").Queries("code", "{code}").HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		client.config.AuthCode = req.FormValue("code")
		_, err := client.config.GetOauthToken()
		if err != nil {
			panic(err)
		}
		server.Shutdown(context.Background())
	})

	server.Handler = router
	server.ListenAndServe()
}
