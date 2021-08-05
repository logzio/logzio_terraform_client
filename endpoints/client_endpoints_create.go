package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	createEndpointServiceUrl    string = endpointServiceEndpoint + "/%s"
	createEndpointServiceMethod string = http.MethodPost
)

// Create a notification endpoint, return the created endpoint's id if successful, an error otherwise
func (c *EndpointsClient) CreateEndpoint(endpoint CreateOrUpdateEndpoint) (*CreateOrUpdateEndpointResponse, error) {
	err := validateCreateOrUpdateEndpointRequest(endpoint)
	if err != nil {
		return nil, err
	}

	createEndpoint, err := json.Marshal(endpoint)
	if err != nil {
		return nil, err
	}

	req, _ := c.buildCreateApiRequest(c.ApiToken, createEndpoint, endpoint.Type)
	jsonResponse, err := logzio_client.CreateHttpRequestBytesResponse(req)
	if err != nil {
		return nil, err
	}

	var retVal CreateOrUpdateEndpointResponse
	err = json.Unmarshal(jsonResponse, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}

func (c *EndpointsClient) buildCreateApiRequest(apiToken string, jsonBytes []byte, endpointType string) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(createEndpointServiceMethod,fmt.Sprintf(createEndpointServiceUrl, baseUrl, c.getURLByType(endpointType)), bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}