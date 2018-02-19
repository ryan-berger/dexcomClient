package dexcomClient

import (
	"encoding/json"
	"fmt"
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

func (client *DexcomClient) GetEvents() []Event {
	req, _ := http.NewRequest("GET",
		urlWithDateRange(client.config, eventsUrl, "2015-09-19T00:00:00", "2015-11-10T00:00:00"), nil)

	//req.Header.Add("authorization", "Bearer "+ client.config.Get)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var response EventResponse
	json.Unmarshal(body, &response)

	return response.Events
}
