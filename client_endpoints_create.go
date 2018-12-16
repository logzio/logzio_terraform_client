package logzio_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const createEndpointServiceUrl string = "%s/v1/endpoints/%s"
const createEndpointServiceMethod string = http.MethodPost
const createEndpointMethodSuccess int = 200

func buildCreateEndpointApiRequest(apiToken string, service string, jsonObject map[string]interface{}) (*http.Request, error) {
	jsonBytes, err := json.Marshal(jsonObject)
	if err != nil {
		return nil, err
	}
	logSomething("buildCreateEndpointApiRequest", fmt.Sprintf("%s", jsonBytes))

	baseUrl := getLogzioBaseUrl()
	req, err := http.NewRequest(createEndpointServiceMethod, fmt.Sprintf(createEndpointServiceUrl, baseUrl, service), bytes.NewBuffer(jsonBytes))
	addHttpHeaders(apiToken, req)

	return req, err
}

func validateCreateEndpointRequest(endpoint EndpointType) (error) {
	return nil
}

func buildCreateEndpointRequest(endpoint EndpointType) map[string]interface{} {
	var createEndpoint = map[string]interface{}{}

	if endpoint.EndpointType == endpointTypeSlack {
		createEndpoint[fldEndpointTitle] = endpoint.Title
		createEndpoint[fldEndpointDescription] = endpoint.Description
		createEndpoint[fldEndpointUrl] = endpoint.Url
	}

	return createEndpoint
}

func (c *Client) createEndpoint(endpoint EndpointType) (*EndpointType, error) {
	err := validateCreateEndpointRequest(endpoint)
	if err != nil {
		return nil, err
	}

	createEndpoint := buildCreateEndpointRequest(endpoint)
	req, _ := buildCreateEndpointApiRequest(c.apiToken, endpoint.EndpointType, createEndpoint)

	var client http.Client
	resp, _ := client.Do(req)
	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	logSomething("CreateEndpoint::Response", fmt.Sprintf("%s", jsonBytes))

	if !checkValidStatus(resp, []int{createEndpointMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "CreateEndpoint", resp.StatusCode, jsonBytes)
	}

	var target EndpointType
	json.Unmarshal(jsonBytes, &target)

	if len(target.Message) > 0 {
		return nil, fmt.Errorf("API call %s failed with status code %d, but with message: %s", "CreateEndpoint", resp.StatusCode, target.Message)
	}

	return &target, nil
}
