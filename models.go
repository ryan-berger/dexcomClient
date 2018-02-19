package dexcomClient


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


type EGVResponse struct {
	Unit  string              `json:"unit"`
	Rate  string              `json:"rate"`
	index int
	EGVS  []*EGV `json:"egvs"`
}

type EGV struct {
	SystemTime  string
	DisplayTime string
	Value       uint64
	Status      string
	Trend       string
	TrendRate   float64
}

type Login struct {
	AccountName string
	Password string
	ApplicationId string
}

type RealTimeData struct {
	DeviceTime string `json:"DT"`
	ServerTime string `json:"ST"`
	Trend int
	Value int
}

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