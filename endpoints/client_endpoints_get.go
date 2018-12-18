package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
	"strings"
)

const getEndpointsServiceUrl string = "%s/v1/endpoints/%d"
const getEndpointsServiceMethod string = http.MethodGet
const getEndpointsMethodSuccess int = 200

func buildGetEnpointApiRequest(apiToken string, notificationId int64) (*http.Request, error) {
	baseUrl := client.GetLogzioBaseUrl()
	req, err := http.NewRequest(getEndpointsServiceMethod, fmt.Sprintf(getEndpointsServiceUrl, baseUrl, notificationId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Returns an endpoint, given it's name.  Returns nil (and an error) if an endpoint with the specified name can't be found
func (c *Endpoints) GetEndpointByName(endpointName string) (*EndpointType, error) {
	list, err := c.ListEndpoints()
	for i := 0; i < len(list); i++ {
		var endpoint EndpointType
		endpoint = list[i]
		if endpoint.Title == endpointName {
			return &endpoint, nil
		}
	}
	return nil, err
}

// Returns an endpoint, given it's identity.  Returns nul (and an error) if an endpoint with the specified id can't be found
func (c *Endpoints) GetEndpoint(endpointId int64) (*EndpointType, error) {
	req, _ := buildGetEnpointApiRequest(c.ApiToken, endpointId)

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{getEndpointsMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "GetEndpoint", resp.StatusCode, jsonBytes)
	}

	str := fmt.Sprintf("%s", jsonBytes)
	if strings.Contains(str, "no endpoint id") {
		return nil, fmt.Errorf("API call %s failed with missing notification %d, data: %s", "GetEndpoint", endpointId, str)
	}

	var jsonEndpoint map[string]interface{}
	err = json.Unmarshal([]byte(jsonBytes), &jsonEndpoint)
	if err != nil {
		return nil, err
	}

	endpoint := jsonEndpointToEndpoint(jsonEndpoint)

	return &endpoint, nil
}
