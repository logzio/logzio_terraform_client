package endpoints

import (
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"net/http"
	"strings"
)

const (
	deleteEndpointServiceUrl    string = endpointServiceEndpoint + "/%d"
	deleteEndpointServiceMethod string = http.MethodDelete
	deleteEndpointMethodSuccess int    = 200
)

const (
	errorDeleteEndpointDoesntExist = "API call DeleteEndpoint failed as endpoint with id:%d doesn't exist, data:%s"
)

func validateDeleteEndpoint(endpoint Endpoint) (error, bool) {
	return nil, true
}

func buildDeleteEndpointApiRequest(apiToken string, service string, endpoint Endpoint) (*http.Request, error) {
	baseUrl := client.GetLogzioBaseUrl()
	req, err := http.NewRequest(deleteEndpointServiceMethod, fmt.Sprintf(deleteEndpointServiceUrl, baseUrl, endpoint.Id), nil)
	logzio_client.AddHttpHeaders(apiToken, req)
	return req, err
}

// Deletes an endpoint with the given id, returns a non nil error otherwise
func (c *Endpoints) DeleteEndpoint(endpointId int64) error {
	if _, err, ok := c.makeEndpointRequest(Endpoint{Id: endpointId}, validateDeleteEndpoint, buildDeleteEndpointApiRequest, func(body []byte) error {
		if strings.Contains(fmt.Sprintf("%s", body), "endpoints/FORBIDDEN_OPERATION") {
			return fmt.Errorf(errorDeleteEndpointDoesntExist, endpointId, body)
		}
		if strings.Contains(fmt.Sprintf("%s", body), "endpoints/UNKNOWN_ENDPOINT") {
			return fmt.Errorf(errorDeleteEndpointDoesntExist, endpointId, body)
		}
		return nil
	}); !ok {
		return err
	}
	return nil
}
