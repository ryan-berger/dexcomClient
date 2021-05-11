package model

// This file contains models for the /egv endpoint. The /evg endpoint
// has a single handler. A sample request/response would look like:
// GET /egv HTTP/1.1
// {
//  "events": [
//    {
//      "systemTime": "2018-02-29T20:59:33",
//      "displayTime": "2018-02-29T12:59:33",
//      "eventType": "carbs",
//      "eventSubType": null,
//      "value": 35,
//      "unit": "grams",
//      "eventId": "601620ec-8caf-4244-81ff-cd946669806e",
//      "eventStatus": "created"
//    },
//    {
//      "systemTime": "2018-02-29T20:59:04",
//      "displayTime": "2018-02-29T12:59:04",
//      "eventType": "insulin",
//      "eventSubType": "longActing",
//      "value": 17,
//      "unit": "units",
//      "eventId": "4b790d7c-542e-4673-b1c4-92e4d8d4550a",
//      "eventStatus": "created"
//    },
//    {
//      "systemTime": "2018-02-29T20:58:53",
//      "displayTime": "2018-02-29T12:58:53",
//      "eventType": "insulin",
//      "eventSubType": "fastActing",
//      "value": 2.5,
//      "unit": "units",
//      "eventId": "a46b5c18-5fbd-4bb4-bcfc-5c9582eede9c",
//      "eventStatus": "created"
//    }
//    {
//      "systemTime": "2018-02-29T20:58:37",
//      "displayTime": "2018-02-29T12:58:37",
//      "eventType": "insulin",
//      "eventSubType": "fastActing",
//      "value": 2,
//      "unit": "units",
//      "eventId": "8705ff4c-e0d4-46a2-ad20-0839eafdfa32",
//      "eventStatus": "deleted"
//    }
//  ]
//}

// EGVResponse is a struct to hold a response from the /egv endpoint
type EGVResponse struct {
	Unit     string `json:"unit"`
	RateUnit string `json:"rateUnit"`
	EGVs     []EGV  `json:"egvs"`
}

// EGV is an Estimated Glucose Value object that contains information
// about a sensor's measurements
type EGV struct {
	RealTimeValue int       `json:"realTimeValue"`
	SmoothedValue int       `json:"smoothedValue"`
	Trend         string    `json:"trend"`
	TrendRate     float64   `json:"trendRate"`
	Time
}
