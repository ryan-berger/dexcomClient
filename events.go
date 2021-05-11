package dexcomClient

import (
	"encoding/json"
	"fmt"
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

func (client *Client) GetEvents(startDate, endDate string) ([]Event, error) {
	req, err := http.NewRequest("GET",
		urlWithDateRange(client.config, eventsUrl, startDate, endDate), nil)


	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "asdf"))
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}


	var response EventResponse
	err = json.
		NewDecoder(resp.Body).
		Decode(&resp)

	if err != nil {
		return nil, err
	}

	return response.Events, nil
}
