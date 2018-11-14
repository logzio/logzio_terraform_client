package logzio_client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const deleteServiceUrl string = "https://api.logz.io/v1/alerts/%d"
const deleteServiceMethod string = "DELETE"

func buildDeleteApiRequest(apiToken string, alertId int64) (*http.Request, error) {
	req, err := http.NewRequest(deleteServiceMethod, fmt.Sprintf(deleteServiceUrl, alertId), nil)
	addHttpHeaders(apiToken, req)
	return req, err
}

func (n *Client) DeleteAlert(alertId int64) error {
	req, _ := buildDeleteApiRequest(n.name, alertId)

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	data, _ := ioutil.ReadAll(resp.Body)

	if !checkValidStatus(resp, []int { 200 }) {
		return fmt.Errorf("%s", data)
	}

	return nil
}