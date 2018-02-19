package dexcomClient

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type DexcomClient struct {
	*http.Client
	*Config
	*AuthClient
	*EstimatedGlucoseClient
}

func NewClient(client *http.Client, config *Config) *DexcomClient {
	dc := &DexcomClient{
		Client:                 client,
		Config:                 config,
		AuthClient:             NewAuthClient(client, config),
		EstimatedGlucoseClient: NewEGVClient(config, &defaultLogger{config: config}),
	}

	if config.IsDev {
		fmt.Println("Dev server starting on :8000")
		fmt.Println(config.GetBaseUrl() + "/v1/oauth2/login?client_id=" + config.ClientId + "&redirect_uri=" + config.RedirectURI + "&response_type=code&scope=offline_access")
		defer dc.startDevServer()
	}
	return dc
}

func NewClientWithAuthCode(client *http.Client, config *Config, authCode string) *DexcomClient {
	config.AuthCode = authCode
	return &DexcomClient{
		Client:                 client,
		Config:                 config,
		AuthClient:             NewAuthClient(client, config),
		EstimatedGlucoseClient: NewEGVClient(config, &defaultLogger{config: config}),
	}
}

func (client *DexcomClient) startDevServer() {
	server := &http.Server{Addr: ":8000"}

	router := mux.NewRouter()
	router.Path("/oauth").Queries("code", "{code}").HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		client.AuthCode = req.FormValue("code")
		_, err := client.GetOauthToken()
		if err != nil {
			panic(err)
		}
		server.Shutdown(context.Background())
	})

	server.Handler = router
	server.ListenAndServe()
}
