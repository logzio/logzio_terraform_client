package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
)

const listEndpointsServiceUrl string = "%s/v1/endpoints"
const listEndpointsServiceMethod string = http.MethodGet
const listEndpointsMethodSuccess int = 200

const errorListEndpointApiCallFailed = "API call ListEndpoints failed with status code:%d, data:%s"

func buildListEndpointsApiRequest(apiToken string) (*http.Request, error) {
	baseUrl := client.GetLogzioBaseUrl()
	req, err := http.NewRequest(listEndpointsServiceMethod, fmt.Sprintf(listEndpointsServiceUrl, baseUrl), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func (c *Endpoints) ListEndpoints() ([]Endpoint, error) {
	req, _ := buildListEndpointsApiRequest(c.ApiToken)

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{listEndpointsMethodSuccess}) {
		return nil, fmt.Errorf(errorListEndpointApiCallFailed, resp.StatusCode, jsonBytes)
	}

	var arr []Endpoint
	var jsonResponse []interface{}
	err = json.Unmarshal([]byte(jsonBytes), &jsonResponse)

	for _, json := range jsonResponse {
		jsonEndpoint := json.(map[string]interface{})
		endpoint := jsonEndpointToEndpoint(jsonEndpoint)
		arr = append(arr, endpoint)
	}

	return arr, nil
}
