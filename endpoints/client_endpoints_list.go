package endpoints

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const listEndpointsServiceUrl string = endpointServiceEndpoint
const listEndpointsServiceMethod string = http.MethodGet
const listEndpointsMethodSuccess int = http.StatusOK

// Returns all the endpoints in an array associated with the account identified by the supplied API token, returns an error if
// any problem occurs during the API call
func (c *EndpointsClient) ListEndpoints() ([]Endpoint, error) {
	req, _ := c.buildListApiRequest(c.ApiToken)
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{listEndpointsMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", listEndpointMethod, resp.StatusCode, jsonBytes)
	}

	var endpoints []Endpoint
	err = json.Unmarshal(jsonBytes, &endpoints)
	if err != nil {
		return nil, err
	}

	return endpoints, nil
}

func (c *EndpointsClient) buildListApiRequest(apiToken string) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(listEndpointsServiceMethod, fmt.Sprintf(listEndpointsServiceUrl, baseUrl), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}
