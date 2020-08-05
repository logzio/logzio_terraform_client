package endpoints

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
	"strings"
)

const getEndpointsServiceUrl string = endpointServiceEndpoint + "/%d"
const getEndpointsServiceMethod string = http.MethodGet
const getEndpointsMethodSuccess int = http.StatusOK

const apiGetEndpointNoEndpoint = "The endpoint doesn't exist"

const errorGetEndpointApiCallFailed = "API call GetEndpoint failed with status code:%d, data:%s"
const errorGetEndpointDoesntExist = "API call GetEndpoint failed as endpoint with id:%d doesn't exist, data:%s"

func (c *EndpointsClient) buildGetEnpointApiRequest(apiToken string, notificationId int64) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(getEndpointsServiceMethod, fmt.Sprintf(getEndpointsServiceUrl, baseUrl, notificationId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Returns an endpoint, given it's name.  Returns nil (and an error) if an endpoint with the specified name can't be found
func (c *EndpointsClient) GetEndpointByName(endpointName string) (*Endpoint, error) {
	list, err := c.ListEndpoints()
	if err != nil {
		return nil, err
	}

	for _, endpoint := range list {
		if endpoint.Title == endpointName {
			return &endpoint, nil
		}
	}

	return nil, err
}

// Returns an endpoint, given it's identity.  Returns nul (and an error) if an endpoint with the specified id can't be found
func (c *EndpointsClient) GetEndpoint(endpointId int64) (*Endpoint, error) {
	req, _ := c.buildGetEnpointApiRequest(c.ApiToken, endpointId)

	jsonEndpoint, err := logzio_client.CreateHttpRequest(req)
	if err != nil {
		return nil, err
	}

	str := fmt.Sprintf("%s", jsonEndpoint)
	if strings.Contains(str, apiGetEndpointNoEndpoint) {
		return nil, fmt.Errorf(errorGetEndpointDoesntExist, endpointId, str)
	}

	endpoint := jsonEndpointToEndpoint(jsonEndpoint)
	return &endpoint, nil
}
