package dexcomClient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	applicationId    = "d89443d2-327c-4a6f-89e5-496bbb0317db"
	agent            = "Dexcom Share/3.0.2.11 CFNetwork/711.2.23 Darwin/14.0.0"
	loginUrl         = "https://share1.dexcom.com/ShareWebServices/Services/General/LoginPublisherAccountByName"
	latestGlucoseUrl = "https://share1.dexcom.com/ShareWebServices/Services/Publisher/ReadPublisherLatestGlucoseValues"
)

type AuthClient struct {
	*http.Client
	*Config
}

func NewAuthClient(httpClient *http.Client, config *Config) *AuthClient {
	return &AuthClient{Client: httpClient, Config: config}
}

func (client *AuthClient) getLatestGlucoseUrl() string {
	return latestGlucoseUrl + "?sessionID=" + client.DexcomToken + "&minutes=1440&maxCount=1"
}

func (client *AuthClient) GetSessionID(username, password string) (string, error) {
	payload := map[string]string{
		"accountName":   username,
		"password":      password,
		"applicationId": applicationId,
	}
	payloadBytes, _ := json.Marshal(&payload)
	payloadReader := bytes.NewReader(payloadBytes)
	req, _ := http.NewRequest("POST", loginUrl, payloadReader)
	req.Header.Add("user-agent", agent)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	sessionId, _ := ioutil.ReadAll(resp.Body)

	var id []byte
	// Strip quotations from the response
	for _, b := range sessionId {
		if b != 34 {
			id = append(id, b)
		}
	}

	client.DexcomToken = string(id)
	return string(id), nil
}

func (client *AuthClient) GetRealTimeData() (*RealTimeData, error) {
	url := client.getLatestGlucoseUrl()

	req, _ := http.NewRequest("POST", url, strings.NewReader(""))
	req.Header.Add("user-agent", agent)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var realTimeData []RealTimeData
	json.Unmarshal(body, &realTimeData)

	if len(realTimeData) == 0 {
		return nil, errors.New("no real time data returned")
	}
	return &realTimeData[0], nil
}

func (client *AuthClient) Authenticate() (*Token, error) {
	return client.GetOauthToken()
}
