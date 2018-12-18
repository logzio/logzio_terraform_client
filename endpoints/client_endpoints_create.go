package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"log"
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

	baseUrl := client.GetLogzioBaseUrl()
	url := fmt.Sprintf(createEndpointServiceUrl, baseUrl, service)
	req, err := http.NewRequest(createEndpointServiceMethod, url, bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	log.Printf("%s, %s, %s", url, req.Header, jsonObject)

	return req, err
}

func validateCreateEndpointRequest(endpoint EndpointType) error {
	return nil
}

func buildCreateEndpointRequest(endpoint EndpointType, service string) map[string]interface{} {
	var createEndpoint = map[string]interface{}{}

	if service == endpointTypeSlack {
		createEndpoint[fldEndpointTitle] = endpoint.Title
		createEndpoint[fldEndpointDescription] = endpoint.Description
		createEndpoint[fldEndpointUrl] = endpoint.Url
	}

	return createEndpoint
}

func (c *Endpoints) createEndpoint(endpoint EndpointType, service string) (*EndpointType, error) {
	err := validateCreateEndpointRequest(endpoint)
	if err != nil {
		return nil, err
	}

	createEndpoint := buildCreateEndpointRequest(endpoint, service)
	req, _ := buildCreateEndpointApiRequest(c.ApiToken, service, createEndpoint)

	var client http.Client
	resp, _ := client.Do(req)
	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{createEndpointMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "CreateEndpoint", resp.StatusCode, jsonBytes)
	}

	var target EndpointType
	json.Unmarshal(jsonBytes, &target)

	if len(target.Message) > 0 {
		return nil, fmt.Errorf("API call %s failed with status code %d, but with message: %s", "CreateEndpoint", resp.StatusCode, target.Message)
	}

	return &target, nil
}
