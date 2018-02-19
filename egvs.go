package dexcomClient

import (
	"io/ioutil"
	"fmt"
	"strconv"
	"sort"
	"net/http"
	"errors"
	"encoding/json"
	"time"
	"math"
)

type rangeRequest struct {
	*Range
	index int
}

type EstimatedGlucoseClient struct {
	*Config
	Logger
}


const EGV_URL = "/v1/users/self/egvs"
const formatString = "2006-01-02"
const dateTimeString = "2006-01-02T15:04:05"

// turns days into nanoseconds
func daysDuration(days int) time.Duration {
	return time.Duration(days) * 24 * time.Hour
}

type Range struct {
	StartDate string
	EndDate string
}

// Get ranges inclusively given a start and end date
func GetEGVRanges(start string, end string) []*Range  {
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

	var ranges []*Range

	for i := 0; i <= monthDistance; i++ {
		var requestEndDay time.Time

		// Start date is 90 * 2i
		requestStartDay := startDate.Add(daysDuration(90 * i + i))
		if i == monthDistance {
			// If we are at the end of the list,
			// don't use the diff formula because then we could
			// be going over the end date
			requestEndDay = endDate
		} else {
			// End date is (90 * 2i) + 90
			// so that we are 90 days ahead of the start date
			requestEndDay = startDate.Add(daysDuration((90 * i + i) + 90))
		}
		ranges = append(ranges,
			&Range{StartDate: requestStartDay.Format(dateTimeString), EndDate: requestEndDay.Format(dateTimeString)})
	}
	return ranges
}




func NewEGVClient(config *Config, logger Logger) *EstimatedGlucoseClient {
	return &EstimatedGlucoseClient{Config: config, Logger: logger}
}

func (client *EstimatedGlucoseClient) GetEGVs(startDate string, endDate string) ([]*EGVResponse, []error) {
	ranges := GetEGVRanges(startDate, endDate)

	fmt.Println(startDate, endDate)

	// egv channel to allow for concurrency in the requests
	egvChan := make(chan *EGVResponse)
	// err chan just in case there are errors
	errChan := make(chan error)

	for i, r := range ranges {
		go client.getRange(rangeRequest{Range: r, index: i}, egvChan, errChan)
	}

	var responses []*EGVResponse
	var errorList []error

	// catch all of the requests. There should only be
	// len(ranges) so we can just wait for all of those
	// to finish
	for i, j := 0, len(ranges); i < j; i++{
		select {
		case response := <- egvChan:
			responses = append(responses, response)
		case reqErr := <- errChan:
			errorList = append(errorList, reqErr)
		}
	}

	sort.Slice(responses, func(i, j int) bool {
		return responses[i].index < responses[j].index
	})

	var egvResponses []*EGVResponse

	for _, response := range responses {
		egvResponses = append(egvResponses, response)
	}

	return egvResponses, errorList
}

func (client *EstimatedGlucoseClient) getRange(request rangeRequest, resultChan chan *EGVResponse, errChan chan error)  {
	// Make request with url date range
	req, _ := http.NewRequest("GET",
		UrlWithDateRange(client.Config, EGV_URL, request.StartDate, request.EndDate), nil)

	client.Debug("URL: " + UrlWithDateRange(client.Config, EGV_URL, request.StartDate, request.EndDate))

	token, err := client.GetOauthToken()

	if err != nil {
		errChan <- err
		return
	}

	req.Header.Add("Authorization", "Bearer " + token.AccessToken)
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
