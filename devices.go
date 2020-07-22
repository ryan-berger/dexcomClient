package dexcomClient

import (
	"encoding/json"
	"io/ioutil"
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

func (client *DexcomClient) GetDevices() ([]*Device, error) {
	url := client.config.getBaseUrl() + "/v1/users/self/devices"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Bearer ")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var devices []*Device

	err = json.Unmarshal(body, &devices)
	if err != nil {
		return nil, err
	}

	return devices, err
}
