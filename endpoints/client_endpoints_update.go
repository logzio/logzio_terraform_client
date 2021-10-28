package endpoints

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	updateEndpointServiceUrl     = endpointServiceEndpoint + "/%s/%d"
	updateEndpointServiceMethod  = http.MethodPut
	updateEndpointMethodSuccess  = http.StatusOK
	updateEndpointMethodNotFound = http.StatusNotFound
)

// UpdateEndpoint updates an existing endpoint, based on the supplied identifier, using the parameters of the specified endpoint
// Returns the updated endpoint id if successful, an error otherwise
func (c *EndpointsClient) UpdateEndpoint(id int64, endpoint CreateOrUpdateEndpoint) (*CreateOrUpdateEndpointResponse, error) {
	err := validateCreateOrUpdateEndpointRequest(endpoint)
	if err != nil {
		return nil, err
	}

	updateEndpointJson, err := json.Marshal(endpoint)
	if err != nil {
		return nil, err
	}

	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   updateEndpointServiceMethod,
		Url:          fmt.Sprintf(updateEndpointServiceUrl, c.BaseUrl, c.getURLByType(endpoint.Type), id),
		Body:         updateEndpointJson,
		SuccessCodes: []int{updateEndpointMethodSuccess},
		NotFoundCode: updateEndpointMethodNotFound,
		ResourceId:   id,
		ApiAction:    updateEndpointMethod,
		ResourceName: endpointResourceName,
	})

	if err != nil {
		return nil, err
	}

	var target CreateOrUpdateEndpointResponse
	err = json.Unmarshal(res, &target)
	if err != nil {
		return nil, err
	}

	return &target, nil
}
