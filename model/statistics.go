package model

import "time"

// StatsRequest is a json struct for all statistical parameters.
type StatsRequest struct {
	Name      string     `json:"name"`
	StartTime time.Time  `json:"startTime"`
	EndTime   time.Time  `json:"endTime"`
	EGVRanges []EGVRange `json:"egvRanges"`
}

// RangeName is a typedef to create an enum of the form:
// enum RangeName { Low, High, UrgentLow }
type RangeName string

const (
	Low       RangeName = "low"
	High      RangeName = "high"
	UrgentLow RangeName = "urgentLow"
)

type EGVRange struct {
	Name  RangeName
	Bound int
}

type StatsResponse struct {
	HypoglycemiaRisk      string  `json:"hypoglycemiaRisk"`
	Min                   int     `json:"min"`
	Max                   int     `json:"max"`
	Mean                  float64 `json:"mean"`
	Median                int     `json:"median"`
	Variance              float64 `json:"variance"`
	StdDev                float64 `json:"stdDev"`
	Sum                   int     `json:"sum"`
	Q1                    int     `json:"q1"`
	Q2                    int     `json:"q2"`
	Q3                    int     `json:"q3"`
	UtilizationPercent    float64 `json:"utilizationPercent"`
	MeanDailyCalibrations float64 `json:"meanDailyCalibrations"`
	NDays                 int     `json:"nDays"`
	NValues               int     `json:"nValues"`
	NUrgentLow            int     `json:"nUrgentLow"`
	NBelowRange           int     `json:"nBelowRange"`
	NWithinRange          int     `json:"nWithinRange"`
}
