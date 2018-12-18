package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
)

const updateEndpointServiceUrl string = "%s/v1/endpoints/%s/%d"
const updateEndpointServiceMethod string = http.MethodPut
const updateEndpointMethodSuccess int = 200

func buildUpdateEndpointApiRequest(apiToken string, id int64, service string, jsonObject map[string]interface{}) (*http.Request, error) {
	jsonBytes, err := json.Marshal(jsonObject)
	if err != nil {
		return nil, err
	}

	baseUrl := client.GetLogzioBaseUrl()
	req, err := http.NewRequest(updateEndpointServiceMethod, fmt.Sprintf(updateEndpointServiceUrl, baseUrl, service, id), bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func validateUpdateEndpointRequest(endpoint EndpointType) error {
	return nil
}

func buildUpdateEndpointRequest(endpoint EndpointType, service string) map[string]interface{} {
	var updateEndpoint = map[string]interface{}{}

	if service == endpointTypeSlack {
		updateEndpoint[fldEndpointTitle] = endpoint.Title
		updateEndpoint[fldEndpointDescription] = endpoint.Description
		updateEndpoint[fldEndpointUrl] = endpoint.Url
	}

	return updateEndpoint
}

func (c *Endpoints) updateEndpoint(id int64, endpoint EndpointType, service string) (*EndpointType, error) {
	err := validateUpdateEndpointRequest(endpoint)
	if err != nil {
		return nil, err
	}

	updateEndpoint := buildUpdateEndpointRequest(endpoint, service)
	req, _ := buildUpdateEndpointApiRequest(c.ApiToken, id, service, updateEndpoint)

	var httpClient http.Client
	resp, _ := httpClient.Do(req)
	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{updateEndpointMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "UpdateEndpoint", resp.StatusCode, jsonBytes)
	}

	var target EndpointType
	json.Unmarshal(jsonBytes, &target)

	if len(target.Message) > 0 {
		return nil, fmt.Errorf("API call %s failed with status code %d, but with message: %s", "UpdateEndpoint", resp.StatusCode, target.Message)
	}

	return &target, nil
}
