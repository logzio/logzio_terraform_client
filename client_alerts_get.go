package logzio_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

const getServiceUrl string = "https://api.logz.io/v1/alerts/%d"
const getServiceMethod string = "GET"

func buildGetApiRequest(apiToken string, alertId int64) (*http.Request, error) {
	req, err := http.NewRequest(getServiceMethod, fmt.Sprintf(getServiceUrl, alertId), nil)
	addHttpHeaders(apiToken, req)
	return req, err
}

func (n *Client) GetAlert(alertId int64) (*AlertType, error) {
	req, _ := buildGetApiRequest(n.name, alertId)

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data, _ := ioutil.ReadAll(resp.Body)

	s, _ := prettyprint(data)
	log.Printf("%s::%s::%s", "some_token", "buildCreateApiRequest", s)

	if !checkValidStatus(resp, []int { 200 }) {
		return nil, fmt.Errorf("%s", data)
	}

	var alert AlertType
	err = json.Unmarshal([]byte(data), &alert)
	if err != nil {
		return nil, err
	}

	return &alert, nil
}