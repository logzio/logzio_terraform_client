package endpoints

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const (
	getEndpointsServiceUrl string = endpointServiceEndpoint + "/%d"
	getEndpointsServiceMethod string = http.MethodGet
	getEndpointsMethodSuccess int = http.StatusOK
	getEndpointsMethodNotFound int = http.StatusNotFound
)

// Returns an endpoint given it's unique identifier, an error otherwise
func (c *EndpointsClient) GetEndpoint(endpointId int64) (*Endpoint, error) {
	req, _ := c.buildGetApiRequest(c.ApiToken, endpointId)
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{getEndpointsMethodSuccess}) {
		if resp.StatusCode == getEndpointsMethodNotFound {
			return nil, fmt.Errorf("API call %s failed with missing endpoint %d, data: %s", getEndpointMethod, endpointId, jsonBytes)
		}

		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", getEndpointMethod, resp.StatusCode, jsonBytes)
	}

	var token Endpoint
	err = json.Unmarshal(jsonBytes, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (c *EndpointsClient) buildGetApiRequest(apiToken string, endpointId int64) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(getEndpointsServiceMethod, fmt.Sprintf(getEndpointsServiceUrl, baseUrl, endpointId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}