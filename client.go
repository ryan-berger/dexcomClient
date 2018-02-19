package dexcomClient

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type DexcomClient struct {
	Logger
	AuthCode    string
	DexcomToken string
	config      *Config
	oAuthToken  *Token
}

func NewClient(config *Config) *DexcomClient {
	dc := &DexcomClient{
		config: config,
	}

	if config.IsDev {
		fmt.Println("Dev server starting on :8000")
		fmt.Println(config.getBaseUrl() + "/v1/oauth2/login?client_id=" + config.ClientId + "&redirect_uri=" + config.RedirectURI + "&response_type=code&scope=offline_access")
		defer dc.startDevServer()
	}
	return dc
}

func NewClientWithToken(config *Config, token *Token) *DexcomClient {
	return &DexcomClient{
		config:     config,
		oAuthToken: token,
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
