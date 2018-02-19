package dexcomClient

import (
	"fmt"
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

func getData() {
	url := "https://sandbox-api.dexcom.com/v1/users/self/devices"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("authorization", "Bearer ")
	resp, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))
}
