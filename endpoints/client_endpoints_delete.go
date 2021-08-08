package endpoints

import (
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const (
	deleteEndpointServiceUrl             string = endpointServiceEndpoint + "/%d"
	deleteEndpointServiceMethod          string = http.MethodDelete
	deleteEndpointMethodSuccess          int    = http.StatusOK
	deleteEndpointMethodSuccessNoContent int    = http.StatusNoContent
	deleteEndpointMethodNotFound         int    = http.StatusNotFound
)

// Delete an endpoint, specified by it's unique id, returns an error if a problem is encountered
func (c *EndpointsClient) DeleteEndpoint(endpointId int64) error {
	req, _ := c.buildDeleteApiRequest(c.ApiToken, endpointId)
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{deleteEndpointMethodSuccess, deleteEndpointMethodSuccessNoContent}) {
		if resp.StatusCode == deleteEndpointMethodNotFound {
			return fmt.Errorf("API call %s failed with missing endpoint %d, data: %s", deleteEndpointMethod, endpointId, jsonBytes)
		}

		return fmt.Errorf("API call %s failed with status code %d, data: %s", deleteEndpointMethod, resp.StatusCode, jsonBytes)
	}

	return nil
}

func (c *EndpointsClient) buildDeleteApiRequest(apiToken string, endpointId int64) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(deleteEndpointServiceMethod, fmt.Sprintf(deleteEndpointServiceUrl, baseUrl, endpointId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}
