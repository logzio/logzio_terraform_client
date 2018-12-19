package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
)

const createEndpointServiceUrl string = "%s/v1/endpoints/%s"
const createEndpointServiceMethod string = http.MethodPost
const createEndpointMethodSuccess int = 200

const errorInvalidEndpointDefinition = "Endpoint definition %s is not valid for service %s"
const errorCreateEndpointApiCallFailed = "API call CreateEndpoint failed with status code %d, data: %s"

func buildCreateEndpointRequest(endpoint EndpointType) map[string]interface{} {
	var createEndpoint = map[string]interface{}{}

	if endpoint.EndpointType == endpointTypeSlack {
		createEndpoint[fldEndpointTitle] = endpoint.Title
		createEndpoint[fldEndpointDescription] = endpoint.Description
		createEndpoint[fldEndpointUrl] = endpoint.Url
	}

	if endpoint.EndpointType == endpointTypeCustom {
		createEndpoint[fldEndpointTitle] = endpoint.Title
		createEndpoint[fldEndpointDescription] = endpoint.Description
		createEndpoint[fldEndpointUrl] = endpoint.Url
		createEndpoint[fldEndpointMethod] = endpoint.Method
		createEndpoint[fldEndpointHeaders] = endpoint.Headers
		createEndpoint[fldEndpointBodyTemplate] = endpoint.BodyTemplate
	}

	if endpoint.EndpointType == endpointTypePagerDuty {
	}

	if endpoint.EndpointType == endpointTypeBigPanda {
	}

	if endpoint.EndpointType == endpointTypeDataDog {
	}

	if endpoint.EndpointType == endpointTypeVictorOps {
	}

	return createEndpoint
}

func buildCreateEndpointApiRequest(apiToken string, service string, jsonObject map[string]interface{}) (*http.Request, error) {
	jsonBytes, err := json.Marshal(jsonObject)
	if err != nil {
		return nil, err
	}

	baseUrl := client.GetLogzioBaseUrl()
	url := fmt.Sprintf(createEndpointServiceUrl, baseUrl, service)
	req, err := http.NewRequest(createEndpointServiceMethod, url, bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Creates an endpoint, given the endpoint definition and the service to create the endpoint against
// Returns the endpoint object if successful (hopefully with an ID) and a non-nil error if not
func (c *Endpoints) CreateEndpoint(endpoint EndpointType) (*EndpointType, error) {
	err := ValidateEndpointRequest(endpoint)
	if err != nil {
		return nil, err
	}

	createEndpoint := buildCreateEndpointRequest(endpoint)
	req, _ := buildCreateEndpointApiRequest(c.ApiToken, endpoint.EndpointType, createEndpoint)

	var httpClient http.Client
	resp, _ := httpClient.Do(req)
	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{createEndpointMethodSuccess}) {
		return nil, fmt.Errorf(errorCreateEndpointApiCallFailed, resp.StatusCode, jsonBytes)
	}

	var target EndpointType
	json.Unmarshal(jsonBytes, &target)

	// logz.io sometimes returns a 200 even though the message is a failure message
	if len(target.Message) > 0 {
		return nil, fmt.Errorf(errorCreateEndpointApiCallFailed, resp.StatusCode, target.Message)
	}

	return &target, nil
}
