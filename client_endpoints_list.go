package logzio_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const listEndpointsServiceUrl string = "%s/v1/endpoints"
const listEndpointsServiceMethod string = http.MethodGet
const listEndpointsMethodSuccess int = 200

func buildListEndpointsApiRequest(apiToken string) (*http.Request, error) {
	baseUrl := getLogzioBaseUrl()
	req, err := http.NewRequest(listEndpointsServiceMethod, fmt.Sprintf(listEndpointsServiceUrl, baseUrl), nil)
	addHttpHeaders(apiToken, req)

	return req, err
}

func (c *Client) ListEndpoints() ([]EndpointType, error) {
	req, _ := buildListEndpointsApiRequest(c.apiToken)

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	logSomething("ListEndpoints::Response", fmt.Sprintf("%s", jsonBytes))

	if !checkValidStatus(resp, []int{listEndpointsMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "ListEndpoints", resp.StatusCode, jsonBytes)
	}

	var arr []EndpointType
	var jsonResponse []interface{}
	err = json.Unmarshal([]byte(jsonBytes), &jsonResponse)

	for x := 0; x < len(jsonResponse); x++ {
		var jsonEndpoint map[string]interface{}
		jsonEndpoint = jsonResponse[x].(map[string]interface{})
		endpoint := jsonEndpointToEndpoint(jsonEndpoint)
		arr = append(arr, endpoint)
	}

	return arr, nil
}
