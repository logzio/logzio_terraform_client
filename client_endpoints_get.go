package logzio_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const getEndpointsServiceUrl string = "%s/v1/endpoints/%d"
const getEndpointsServiceMethod string = http.MethodGet
const getEndpointsMethodSuccess int = 200

func buildGetEnpointApiRequest(apiToken string, notificationId int64) (*http.Request, error) {
	baseUrl := getLogzioBaseUrl()
	req, err := http.NewRequest(getEndpointsServiceMethod, fmt.Sprintf(getEndpointsServiceUrl, baseUrl, notificationId), nil)
	addHttpHeaders(apiToken, req)

	return req, err
}

func (c *Client) GetEndpoint(endpointId int64) (*EndpointType, error) {
	req, _ := buildGetEnpointApiRequest(c.apiToken, endpointId)

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	logSomething("GetEndpoint::Response", fmt.Sprintf("%s", jsonBytes))

	if !checkValidStatus(resp, []int{getEndpointsMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "GetEndpoint", resp.StatusCode, jsonBytes)
	}

	str := fmt.Sprintf("%s", jsonBytes)
	if strings.Contains(str, "no endpoint id") {
		return nil, fmt.Errorf("API call %s failed with missing notification %d, data: %s", "GetEndpoint", endpointId, str)
	}

	var jsonEndpoint map[string]interface{}
	err = json.Unmarshal([]byte(jsonBytes), &jsonEndpoint)
	if err != nil {
		return nil, err
	}

	endpoint := jsonEndpointToEndpoint(jsonEndpoint)

	return &endpoint, nil
}