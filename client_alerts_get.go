package logzio_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

const getServiceUrl string = "%s/v1/alerts/%d"
const getServiceMethod string = "GET"

func buildGetApiRequest(apiToken string, alertId int64) (*http.Request, error) {

	baseUrl := getLogzioBaseUrl()
	req, err := http.NewRequest(getServiceMethod, fmt.Sprintf(getServiceUrl, baseUrl, alertId), nil)

	addHttpHeaders(apiToken, req)
	return req, err
}

func (c *Client) GetAlert(alertId int64) (*AlertType, error) {
	req, _ := buildGetApiRequest(c.apiToken, alertId)

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	s, _ := prettyprint(data)
	log.Printf("%s::%s", "GetAlert", s)

	if !checkValidStatus(resp, []int { 200 }) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "GetAlert", resp.StatusCode, s)
	}

	var alert AlertType
	err = json.Unmarshal([]byte(data), &alert)
	if err != nil {
		return nil, err
	}

	return &alert, nil
}