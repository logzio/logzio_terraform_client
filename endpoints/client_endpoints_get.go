package endpoints

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	getEndpointsServiceUrl     = endpointServiceEndpoint + "/%d"
	getEndpointsServiceMethod  = http.MethodGet
	getEndpointsMethodSuccess  = http.StatusOK
	getEndpointsMethodNotFound = http.StatusNotFound
)

// GetEndpoint returns an endpoint given its unique identifier, an error otherwise
func (c *EndpointsClient) GetEndpoint(endpointId int64) (*Endpoint, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   getEndpointsServiceMethod,
		Url:          fmt.Sprintf(getEndpointsServiceUrl, c.BaseUrl, endpointId),
		Body:         nil,
		SuccessCodes: []int{getEndpointsMethodSuccess},
		NotFoundCode: getEndpointsMethodNotFound,
		ResourceId:   endpointId,
		ApiAction:    getEndpointMethod,
	})

	if err != nil {
		return nil, err
	}

	var endpoint Endpoint
	err = json.Unmarshal(res, &endpoint)
	if err != nil {
		return nil, err
	}

	return &endpoint, nil
}
