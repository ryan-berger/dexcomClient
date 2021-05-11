package dexcomClient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"sort"
	"strconv"
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

type Range struct {
	StartDate string
	EndDate   string
}

func (client *Client) GetEGVs(startDate string, endDate string) ([]*EGVResponse, error) {
	ranges := getEGVRanges(startDate, endDate)

	// egv channel to allow for concurrency in the requests
	egvChan := make(chan *EGVResponse)
	// err chan just in case there are errors
	errChan := make(chan error)

	for i, r := range ranges {
		go client.getRange(rangeRequest{Range: r, index: i}, egvChan, errChan)
	}

	var responses []*EGVResponse

	// catch all of the requests. There should only be
	// len(ranges) so we can just wait for all of those
	// to finish
	for i, j := 0, len(ranges); i < j; i++ {
		select {
		case response := <-egvChan:
			responses = append(responses, response)
		case reqErr := <-errChan:
			return nil, reqErr
		}
	}

	sort.Slice(responses, func(i, j int) bool {
		return responses[i].index < responses[j].index
	})

	var egvResponses []*EGVResponse

	for _, response := range responses {
		egvResponses = append(egvResponses, response)
	}

	return egvResponses, nil
}

func (client *Client) getRange(request rangeRequest, resultChan chan *EGVResponse, errChan chan error) {
	// Make request with url date range
	req, _ := http.NewRequest("GET",
		urlWithDateRange(client.config, egvUrl, request.StartDate, request.EndDate), nil)

	client.Debug("URL: " + urlWithDateRange(client.config, egvUrl, request.StartDate, request.EndDate))


	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errChan <- err
		panic(err)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		fmt.Println(string(body))
		panic(errors.New("Code " + strconv.FormatInt(int64(resp.StatusCode), 10)))
	}

	var jsonResponse EGVResponse
	jsonErr := json.Unmarshal(body, &jsonResponse)

	// deal with json parse error
	if jsonErr != nil {
		errChan <- jsonErr
		panic(jsonErr)
		return
	}
	jsonResponse.index = request.index
	resultChan <- &jsonResponse
}

type rangeRequest struct {
	Range
	index int
}

// turns days into nanoseconds
func daysDuration(days int) time.Duration {
	return time.Duration(days) * 24 * time.Hour
}

// Get ranges inclusively given a start and end date
func getEGVRanges(start string, end string) []Range {
	// Parse dates
	startDate, _ := time.Parse(dateTimeString, start)
	endDate, _ := time.Parse(dateTimeString, end)

	endUTC := endDate.UTC()

	// Get diff between the dates in days
	diff := endUTC.Sub(startDate)
	daysDiff := diff.Hours() / 24

	// Figure out how many months we have to request
	// and subtract one because it may not be completely even
	monthDistance := int(math.Floor(daysDiff / 90))

	var ranges []Range

	for i := 0; i <= monthDistance; i++ {
		var requestEndDay time.Time

		// Start date is 90 * 2i
		requestStartDay := startDate.Add(daysDuration(90*i + i))
		if i == monthDistance {
			// If we are at the end of the list,
			// don't use the diff formula because then we could
			// be going over the end date
			requestEndDay = endDate
		} else {
			// End date is (90 * 2i) + 90
			// so that we are 90 days ahead of the start date
			requestEndDay = startDate.Add(daysDuration((90*i + i) + 90))
		}
		ranges = append(ranges,
			Range{StartDate: requestStartDay.Format(dateTimeString), EndDate: requestEndDay.Format(dateTimeString)})
	}
	return ranges
}
