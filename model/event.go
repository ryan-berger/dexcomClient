package model

// This file holds all models for the /events endpoint, which has one handler.
// /
// The response represents the JSON object:
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
// }

// EventResponse holds an array of events an can be fetched
// via the /events
type EventResponse struct {
	Events []Event `json:"events"`
}

// Event is a struct to hold arbitrary Dexcom user events such as
// exercise events,
type Event struct {
	ID      string  `json:"id"`
	Type    string  `json:"eventType"`
	SubType *string `json:"eventSubType"`
	Value   int     `json:"value"`
	Unit    string  `json:"unit"`
	Status  string  `json:"eventStatus"`
	Time
}
