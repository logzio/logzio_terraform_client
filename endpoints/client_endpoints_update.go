package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const (
	updateEndpointServiceUrl     string = endpointServiceEndpoint + "/%s/%d"
	updateEndpointServiceMethod  string = http.MethodPut
	updateEndpointMethodSuccess  int    = http.StatusOK
	updateEndpointMethodNotFound int    = http.StatusNotFound
)

// Updates an existing endpoint, based on the supplied identifier, using the parameters of the specified endpoint
// Returns the updated endpoint id if successful, an error otherwise
func (c *EndpointsClient) UpdateEndpoint(id int64, endpoint CreateOrUpdateEndpoint) (*CreateOrUpdateEndpointResponse, error) {
	err := validateCreateOrUpdateEndpointRequest(endpoint)
	if err != nil {
		return nil, err
	}

	req, err := c.buildUpdateEndpointApiRequest(c.ApiToken, endpoint, id)
	if err != nil {
		return nil, err
	}

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{updateEndpointMethodSuccess}) {
		if resp.StatusCode == updateEndpointMethodNotFound {
			return nil, fmt.Errorf("API call %s failed with missing endpoint %d, data: %s", updateEndpointMethod, id, jsonBytes)
		}

		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", updateEndpointMethod, resp.StatusCode, jsonBytes)
	}

	var target CreateOrUpdateEndpointResponse
	err = json.Unmarshal(jsonBytes, &target)
	if err != nil {
		return nil, err
	}

	return &target, nil
}

func (c *EndpointsClient) buildUpdateEndpointApiRequest(apiToken string, endpoint CreateOrUpdateEndpoint, endpointId int64) (*http.Request, error) {
	jsonBytes, err := json.Marshal(endpoint)
	if err != nil {
		return nil, err
	}

	baseUrl := c.BaseUrl
	req, err := http.NewRequest(updateEndpointServiceMethod, fmt.Sprintf(updateEndpointServiceUrl, baseUrl, c.getURLByType(endpoint.Type), endpointId), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}
