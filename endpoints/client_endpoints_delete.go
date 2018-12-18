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

const deleteEndpointFnId = "DeleteEndpoint"

func buildDeleteEndpointApiRequest(apiToken string, endpointId int64) (*http.Request, error) {
	baseUrl := client.GetLogzioBaseUrl()
	req, err := http.NewRequest(deleteEndpointServiceMethod, fmt.Sprintf(deleteEndpointServiceUrl, baseUrl, endpointId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func (c *Endpoints) DeleteEndpoint(endpointId int64) error {
	req, _ := buildDeleteEndpointApiRequest(c.ApiToken, endpointId)

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{deleteEndpointMethodSuccess}) {
		return fmt.Errorf("API call %s failed with status code %d, data: %s", deleteEndpointFnId, resp.StatusCode, jsonBytes)
	}

	str := fmt.Sprintf("%s", jsonBytes)
	if strings.Contains(str, "no endpoint id") {
		return fmt.Errorf("API call %s failed with missing endpoint %d, data: %s", deleteEndpointFnId, endpointId, jsonBytes)
	}

	return nil
}
