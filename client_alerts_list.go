package logzio_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const listServiceUrl string = "https://api.logz.io/v1/alerts"
const listServiceMethod string = "GET"

func buildListApiRequest(apiToken string) (*http.Request, error) {
	req, err := http.NewRequest(listServiceMethod, listServiceUrl, nil)
	addHttpHeaders(apiToken, req)
	return req, err
}

func (n *Client) ListAlerts() ([]AlertType, error) {
	req, _ := buildListApiRequest(n.name)

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data, _ := ioutil.ReadAll(resp.Body)

	if !checkValidStatus(resp, []int{200}) {
		return nil, fmt.Errorf("%s", data)
	}

	var arr []AlertType
	err = json.Unmarshal([]byte(data), &arr)
	if err != nil {
		return nil, err
	}

	return arr, nil
}
