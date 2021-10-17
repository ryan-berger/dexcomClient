package dexcomClient

import (
	"encoding/json"
	"net/http"
	"sort"
	"time"
)

const egvUrl = "/v1/users/self/egvs"
const dateTimeString = "2006-01-02T15:04:05"

type EGVResponse struct {
	Unit  string `json:"unit"`
	Rate  string `json:"rate"`
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

type queryParam struct {
	StartDate string
	EndDate   string
}

func (c *Client) GetEGVs(startDate string, endDate string) ([]EGVResponse, error) {
	ranges, err := getEGVRanges(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// egv channel to allow for concurrency in the requests
	egvChan := make(chan *EGVResponse)
	// err chan just in case there are errors
	errChan := make(chan error)

	for i, r := range ranges {
		go c.getRange(rangeRequest{queryParam: r, index: i}, egvChan, errChan)
	}

	var responses []EGVResponse

	// catch all of the requests. There should only be
	// len(ranges) so we can just wait for all of those
	// to finish
	for i, j := 0, len(ranges); i < j; i++ {
		select {
		case response := <-egvChan:
			responses = append(responses, *response)
		case reqErr := <-errChan:
			return nil, reqErr
		}
	}

	sort.Slice(responses, func(i, j int) bool {
		return responses[i].index < responses[j].index
	})

	var egvResponses []EGVResponse

	for _, response := range responses {
		egvResponses = append(egvResponses, response)
	}

	return egvResponses, nil
}

func (c *Client) getRange(request rangeRequest, resultChan chan *EGVResponse, errChan chan error) {
	// Make request with url date range
	req, _ := http.NewRequest("GET", "", nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errChan <- err
		return
	}

	var jsonResponse EGVResponse
	jsonErr := json.
		NewDecoder(resp.Body).
		Decode(&jsonResponse)

	// deal with json parse error
	if jsonErr != nil {
		errChan <- jsonErr
		return
	}

	jsonResponse.index = request.index
	resultChan <- &jsonResponse
}

type rangeRequest struct {
	queryParam
	index int
}

// turns days into nanoseconds
func daysDuration(days int) time.Duration {
	return time.Duration(days) * 24 * time.Hour
}

func offset(t time.Time, days int) string {
	return t.
		Add(daysDuration(days)).
		Format(dateTimeString)
}

// Get ranges inclusively given a start and end date
func getEGVRanges(start string, end string) ([]queryParam, error) {
	// Parse dates
	var startDate, endDate time.Time
	var err error

	startDate, err = time.Parse(dateTimeString, start)
	if err != nil {
		return nil, err
	}

	endDate, err = time.Parse(dateTimeString, end)
	if err != nil {
		return nil, err
	}

	endUTC := endDate.UTC()

	// Get diff between the dates in days
	diff := endUTC.Sub(startDate.UTC())
	days := int(diff.Hours() / 24)

	var ranges []queryParam

	for i := 0; i < days / 90; i++ {
		ranges = append(
			ranges,
			queryParam{
				StartDate: offset(startDate, i * 90),
				EndDate: offset(startDate, (i + 1) * 90),
			},
		)
	}

	if leftover := days % 90; leftover != 0 {
		startOffset := days - leftover
		ranges = append(
			ranges,
			queryParam{
				StartDate: offset(startDate, startOffset),
				EndDate: offset(startDate, days),
			},
		)
	}

	return ranges, nil
}
