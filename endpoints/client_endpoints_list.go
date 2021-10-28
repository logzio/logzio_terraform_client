package endpoints

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"net/http"
)

const (
	listEndpointsServiceUrl     = endpointServiceEndpoint
	listEndpointsServiceMethod  = http.MethodGet
	listEndpointsMethodSuccess  = http.StatusOK
	listEndpointsStatusNotFound = http.StatusNotFound
)

// ListEndpoints returns all the endpoints in an array associated with the account identified by the supplied API token, returns an error if
// any problem occurs during the API call
func (c *EndpointsClient) ListEndpoints() ([]Endpoint, error) {
	res, err := logzio_client.CallLogzioApi(logzio_client.LogzioApiCallDetails{
		ApiToken:     c.ApiToken,
		HttpMethod:   listEndpointsServiceMethod,
		Url:          fmt.Sprintf(listEndpointsServiceUrl, c.BaseUrl),
		Body:         nil,
		SuccessCodes: []int{listEndpointsMethodSuccess},
		NotFoundCode: listEndpointsStatusNotFound,
		ResourceId:   nil,
		ApiAction:    listEndpointMethod,
		ResourceName: endpointResourceName,
	})

	if err != nil {
		return nil, err
	}

	var endpoints []Endpoint
	err = json.Unmarshal(res, &endpoints)
	if err != nil {
		return nil, err
	}

	return endpoints, nil
}
