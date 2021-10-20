package endpoints

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	createEndpointServiceUrl      = endpointServiceEndpoint + "/%s"
	createEndpointServiceMethod   = http.MethodPost
	createEndpointMethodSuccess   = http.StatusOK
	createEndpointMethodNoContent = http.StatusNoContent
	createEndpointMethodCreated   = http.StatusCreated
	createEndpointStatusNotFound  = http.StatusNotFound
)

// CreateEndpoint creates a notification endpoint, return the created endpoint's id if successful, an error otherwise
func (c *EndpointsClient) CreateEndpoint(endpoint CreateOrUpdateEndpoint) (*CreateOrUpdateEndpointResponse, error) {
	err := validateCreateOrUpdateEndpointRequest(endpoint)
	if err != nil {
		return nil, err
	}

	createEndpointJson, err := json.Marshal(endpoint)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   createEndpointServiceMethod,
		Url:          fmt.Sprintf(createEndpointServiceUrl, c.BaseUrl, c.getURLByType(endpoint.Type)),
		Body:         createEndpointJson,
		SuccessCodes: []int{createEndpointMethodSuccess, createEndpointMethodNoContent, createEndpointMethodCreated},
		NotFoundCode: createEndpointStatusNotFound,
		ResourceId:   nil,
		ApiAction:    createEndpointMethod,
	})

	if err != nil {
		return nil, err
	}

	var retVal CreateOrUpdateEndpointResponse
	err = json.Unmarshal(res, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil
}
