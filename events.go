package dexcomClient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const eventsUrl = "v1/users/self/events"

type Event struct {
	SystemTime   string `json:"systemTime"`
	DisplayTime  string `json:"displayTime"`
	EventType    string `json:"eventType"`
	EventSubType string `json:"eventSubType"`
	Value        int    `json:"value"`
	Unit         string `json:"unit"`
}

type EventResponse struct {
	Events []Event `json:"events"`
}

func (client *DexcomClient) GetEvents(startDate, endDate string) ([]Event, error) {
	req, _ := http.NewRequest("GET",
		urlWithDateRange(client.config, eventsUrl, startDate, endDate), nil)

	token, err := client.GetOauthToken()

	if err != nil {
		return nil, err
	}

	req.Header.Add("authorization", "Bearer " + token.AccessToken)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var response EventResponse
	json.Unmarshal(body, &response)

	return response.Events, nil
}
