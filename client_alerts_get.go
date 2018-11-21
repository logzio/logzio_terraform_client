package logzio_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const getServiceUrl string = "%s/v1/alerts/%d"
const getServiceMethod string = http.MethodGet
const getMethodSuccess int = 200

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

	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	logSomething("GetAlert::Response", fmt.Sprintf("%s", jsonBytes))

	if !checkValidStatus(resp, []int{getMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "GetAlert", resp.StatusCode, jsonBytes)
	}

	str := fmt.Sprintf("%s", jsonBytes)
	if strings.Contains(str, "no alert id") {
		return nil, fmt.Errorf("API call %s failed with missing alert %d, data: %s", "GetAlert", alertId, str)
	}

	var jsonAlert map[string]interface{}
	err = json.Unmarshal([]byte(jsonBytes), &jsonAlert)
	if err != nil {
		return nil, err
	}

	alert := jsonAlertToAlert(jsonAlert)

	return &alert, nil
}
