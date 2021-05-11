package model

import "time"


// DataRange is an object which can inform further requests and
// valid date ranges for requesting user data
type DataRange struct {
	Calibrations Range
	EVGs         Range
	Events       []Event
}

// Time is a struct that handles both systemTime and displayTime which
// is used frequently throughout the Dexcom API
type Time struct {
	SystemTime  time.Time `json:"systemTime"`
	DisplayTime time.Time `json:"displayTime"`
}

// Range is a struct that holds a start and end Time
type Range struct {
	Start Time
	End   Time
}