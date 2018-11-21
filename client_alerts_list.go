package logzio_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const listServiceUrl string = "%s/v1/alerts"
const listServiceMethod string = http.MethodGet
const listMethodSuccess int = 200

func buildListApiRequest(apiToken string) (*http.Request, error) {
	baseUrl := getLogzioBaseUrl()
	req, err := http.NewRequest(listServiceMethod, fmt.Sprintf(listServiceUrl, baseUrl), nil)
	addHttpHeaders(apiToken, req)

	return req, err
}

func (c *Client) ListAlerts() ([]AlertType, error) {
	req, _ := buildListApiRequest(c.apiToken)

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	logSomething("ListAlerts::Response", fmt.Sprintf("%s", jsonBytes))

	if !checkValidStatus(resp, []int{listMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "ListAlerts", resp.StatusCode, jsonBytes)
	}

	var arr []AlertType
	var jsonResponse []interface{}
	err = json.Unmarshal([]byte(jsonBytes), &jsonResponse)

	for x := 0; x < len(jsonResponse); x++ {
		var jsonAlert map[string]interface{}
		jsonAlert = jsonResponse[x].(map[string]interface{})

		alert := jsonAlertToAlert(jsonAlert)

		arr = append(arr, alert)
	}

	return arr, nil
}