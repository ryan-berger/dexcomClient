package dexcomClient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Device struct {
	Model          string
	LastUploadDate string
	AlertSettings  []AlertSetting
}

type AlertSetting struct {
	AlertName   string
	Value       int
	Unit        string
	Snooze      int
	Delay       int
	Enabled     bool
	SystemTime  string
	DisplayTime string
}

func (client *Client) GetDevices() ([]Device, error) {
	url := fmt.Sprintf("%s/v1/users/self/devices", client.config.getBaseUrl())
	req, err := http.NewRequest("GET", url, nil)

	// TODO:
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer ")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var devices []Device
	err = json.
		NewDecoder(resp.Body).
		Decode(&devices)

	if err != nil {
		return nil, err
	}

	return devices, nil
}
