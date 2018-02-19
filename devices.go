package dexcomClient

import (
	"io/ioutil"
	"fmt"
	"net/http"
)

func GetData() {
	url := "https://sandbox-api.dexcom.com/v1/users/self/devices"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("authorization", "Bearer ")
	resp, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))
}
