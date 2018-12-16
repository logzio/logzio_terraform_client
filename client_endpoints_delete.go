package logzio_client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const deleteEndpointServiceUrl string = "%s/v1/endpoints/%d"
const deleteEndpointServiceMethod string = http.MethodDelete
const deleteEndpointMethodSuccess int = 200

const deleteEndpointFnId = "DeleteEndpoint"

func buildDeleteEndpointApiRequest(apiToken string, endpointId int64) (*http.Request, error) {
	baseUrl := getLogzioBaseUrl()
	req, err := http.NewRequest(deleteEndpointServiceMethod, fmt.Sprintf(deleteEndpointServiceUrl, baseUrl, endpointId), nil)
	addHttpHeaders(apiToken, req)
	logSomething("buildDeleteEndpointApiRequest", req.URL.Path)

	return req, err
}

func (c *Client) DeleteEndpoint(endpointId int64) error {
	req, _ := buildDeleteEndpointApiRequest(c.apiToken, endpointId)

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	logSomething(deleteEndpointFnId + "::Response", fmt.Sprintf("%s", jsonBytes))

	if !checkValidStatus(resp, []int{deleteEndpointMethodSuccess}) {
		return fmt.Errorf("API call %s failed with status code %d, data: %s", deleteEndpointFnId, resp.StatusCode, jsonBytes)
	}

	str := fmt.Sprintf("%s", jsonBytes)
	if strings.Contains(str, "no endpoint id") {
		return fmt.Errorf("API call %s failed with missing endpoint %d, data: %s", deleteEndpointFnId, endpointId, jsonBytes)
	}

	return nil
}
