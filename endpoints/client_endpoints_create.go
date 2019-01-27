package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
	"strings"
)

const createEndpointServiceUrl string = "%s/v1/endpoints/%s"
const createEndpointServiceMethod string = http.MethodPost
const createEndpointMethodSuccess int = 200

const errorInvalidEndpointDefinition = "endpoint definition %v is not valid for service %s"
const errorCreateEndpointApiCallFailed = "API call CreateEndpoint failed with status code %d, data: %s"

func buildCreateEndpointRequest(endpoint Endpoint) map[string]interface{} {
	var createEndpoint = map[string]interface{}{}

	createEndpoint[fldEndpointTitle] = endpoint.Title
	createEndpoint[fldEndpointDescription] = endpoint.Description

	if endpoint.EndpointType == EndpointTypeSlack {
		createEndpoint[fldEndpointUrl] = endpoint.Url
	}

	if endpoint.EndpointType == EndpointTypeCustom {
		createEndpoint[fldEndpointUrl] = endpoint.Url
		createEndpoint[fldEndpointMethod] = endpoint.Method
		headers := endpoint.Headers
		headerStrings := []string{}
		for k, v := range headers {
			headerStrings = append(headerStrings, fmt.Sprintf("%s=%s", k, v))
		}
		headerString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(headerStrings)), ","), "[]")
		createEndpoint[fldEndpointHeaders] = headerString
		createEndpoint[fldEndpointBodyTemplate] = endpoint.BodyTemplate
	}

	if endpoint.EndpointType == EndpointTypePagerDuty {
		createEndpoint[fldEndpointServiceKey] = endpoint.ServiceKey
	}

	if endpoint.EndpointType == EndpointTypeBigPanda {
		createEndpoint[fldEndpointApiToken] = endpoint.ApiToken
		createEndpoint[fldEndpointAppKey] = endpoint.AppKey
	}

	if endpoint.EndpointType == EndpointTypeDataDog {
		createEndpoint[fldEndpointApiKey] = endpoint.ApiKey
	}

	if endpoint.EndpointType == EndpointTypeVictorOps {
		createEndpoint[fldEndpointRoutingKey] = endpoint.RoutingKey
		createEndpoint[fldEndpointMessageType] = endpoint.MessageType
		createEndpoint[fldEndpointServiceApiKey] = endpoint.ServiceApiKey
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
func (c *Endpoints) CreateEndpoint(endpoint Endpoint) (*Endpoint, error) {
	err := ValidateEndpointRequest(endpoint)
	if err != nil {
		return nil, err
	}

	createEndpoint := buildCreateEndpointRequest(endpoint)
	req, _ := buildCreateEndpointApiRequest(c.ApiToken, endpoint.EndpointType, createEndpoint)

	httpClient := client.GetHttpClient(req)

	resp, _ := httpClient.Do(req)
	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{createEndpointMethodSuccess}) {
		return nil, fmt.Errorf(errorCreateEndpointApiCallFailed, resp.StatusCode, jsonBytes)
	}

	var target Endpoint
	json.Unmarshal(jsonBytes, &target)

	// logz.io sometimes returns a 200 even though the message is a failure message
	if len(target.Message) > 0 {
		return nil, fmt.Errorf(errorCreateEndpointApiCallFailed, resp.StatusCode, target.Message)
	}

	return &target, nil
}
