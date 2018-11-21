package logzio_client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const deleteServiceUrl string = "%s/v1/alerts/%d"
const deleteServiceMethod string = http.MethodDelete
const deleteMethodSuccess int = 200

func buildDeleteApiRequest(apiToken string, alertId int64) (*http.Request, error) {
	baseUrl := getLogzioBaseUrl()
	req, err := http.NewRequest(deleteServiceMethod, fmt.Sprintf(deleteServiceUrl, baseUrl, alertId), nil)
	addHttpHeaders(apiToken, req)
	logSomething("buildDeleteApiRequest", req.URL.Path)

	return req, err
}

func (c *Client) DeleteAlert(alertId int64) error {
	req, _ := buildDeleteApiRequest(c.apiToken, alertId)

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	logSomething("DeleteAlert::Response", fmt.Sprintf("%s", jsonBytes))

	if !checkValidStatus(resp, []int{deleteMethodSuccess}) {
		return fmt.Errorf("API call %s failed with status code %d, data: %s", "DeleteAlert", resp.StatusCode, jsonBytes)
	}

	str := fmt.Sprintf("%s", jsonBytes)
	if strings.Contains(str, "no alert id") {
		return fmt.Errorf("API call %s failed with missing alert %d, data: %s", "DeleteAlert", alertId, jsonBytes)
	}

	return nil
}
