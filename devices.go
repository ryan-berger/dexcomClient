package dexcomClient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ryan-berger/dexcomClient/model"
)


type deviceResponse struct {
	Devices []model.Device `json:"devices"`
}

func (c *Client) GetDevices() ([]model.Device, error) {
	url := c.getURL("/v2/users/self/devices")
	req, err := http.NewRequest("GET", url, nil)

	// TODO:
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var devResp deviceResponse
	err = json.
		NewDecoder(resp.Body).
		Decode(&devResp)

	if err != nil {
		return nil, err
	}

	return devResp.Devices, nil
}
