package model

type Device struct {
	Model          string `json:"model"`
	LastUploadDate string `json:"lastUploadDate"`
	AlertSettings  []AlertSetting
}

type AlertSetting struct {
	AlertName string `json:"alertName"`
	Value     int    `json:"value"`
	Unit      string `json:"unit"`
	Snooze    int    `json:"snooze"`
	Delay     int    `json:"delay"`
	Enabled   bool   `json:"enabled"`
	Time
}

type DeviceResponse struct {
	Devices []Device `json:"devices"`
}
