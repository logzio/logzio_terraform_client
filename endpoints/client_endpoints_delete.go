package endpoints

import (
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
	"strings"
)

const deleteEndpointServiceUrl string = "%s/v1/endpoints/%d"
const deleteEndpointServiceMethod string = http.MethodDelete
const deleteEndpointMethodSuccess int = 200

const errorDeleteEndpointApiCallFailed = "API call DeleteEndpoint failed with status code:%d, data:%s"
const errorDeleteEndpointDoesntExist = "API call DeleteEndpoint failed as endpoint with id:%d doesn't exist, data:%s"

const apiDeleteEndpointNoEndpoint = "no endpoint id"

func buildDeleteEndpointApiRequest(apiToken string, endpointId int64) (*http.Request, error) {
	baseUrl := client.GetLogzioBaseUrl()
	req, err := http.NewRequest(deleteEndpointServiceMethod, fmt.Sprintf(deleteEndpointServiceUrl, baseUrl, endpointId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)
	return req, err
}

// Deletes an endpoint with the given id, returns a non nil error otherwise
func (c *Endpoints) DeleteEndpoint(endpointId int64) error {
	req, _ := buildDeleteEndpointApiRequest(c.ApiToken, endpointId)
	httpClient := client.GetHttpClient(req)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{deleteEndpointMethodSuccess}) {
		return fmt.Errorf(errorDeleteEndpointApiCallFailed, resp.StatusCode, jsonBytes)
	}

	str := fmt.Sprintf("%s", jsonBytes)
	if strings.Contains(str, apiDeleteEndpointNoEndpoint) {
		return fmt.Errorf(errorDeleteEndpointDoesntExist, endpointId, jsonBytes)
	}

	return nil
}
