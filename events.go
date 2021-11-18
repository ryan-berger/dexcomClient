package dexcomClient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ryan-berger/dexcomClient/model"
)

type eventResponse struct {
	Events []model.Event `json:"events"`
}

func (c *Client) GetEvents(startDate, endDate string) ([]model.Event, error) {
	path := fmt.Sprintf("/v2/users/self/egvs?startDate=%s&endDate=%s", startDate, endDate)

	req, err := http.NewRequest("GET", c.getURL(path), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	var response eventResponse
	err = json.
		NewDecoder(resp.Body).
		Decode(&resp)

	if err != nil {
		return nil, err
	}

	return response.Events, nil
}
