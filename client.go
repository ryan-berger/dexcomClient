package dexcomClient

import (
	"golang.org/x/oauth2"
)

type Client struct {
	AuthCode    string
	DexcomToken string
	config      *Config
	conf        oauth2.Config
	logger
}

func NewClient(config *Config) *Client {
	dc := &Client{
		config: config,
		logger: &defaultLogger{config: config},
	}
	return dc
}