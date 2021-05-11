package model

import "time"


// CalibrationResponse is a struct that holds information about
// user calibration data. It can be retrieved from the
// /calibrations endpoint. From the docs, the JSON looks like:
// {
//  "calibrations": [
//    {
//      "systemTime": "2017-06-17T03:59:11",
//      "displayTime": "2017-06-16T19:59:11",
//      "unit": "mg/dL",
//      "value": 124
//    },
//    {
//      "systemTime": "2017-06-16T16:03:28",
//      "displayTime": "2017-06-16T08:03:28",
//      "unit": "mg/dL",
//      "value": 86
//    }
//  ]
//}
type CalibrationResponse struct {
	Calibrations []Calibration `json:"calibrations"`
}

// Calibration represents a single calibration event for a Dexcom sensor
type Calibration struct {
	SystemTime  time.Time `json:"systemTime"`
	DisplayTime time.Time `json:"displayTime"`
	Unit        string    `json:"unit"`
	Value       int       `json:"value"`
}
