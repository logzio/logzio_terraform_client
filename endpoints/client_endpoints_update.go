package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"net/http"
	"strings"
)

const (
	updateEndpointServiceUrl string = endpointServiceEndpoint + "/%s/%d"
	updateEndpointServiceMethod string = http.MethodPut
	updateEndpointMethodSuccess int = 200
)

const (
	errorUpdateEndpointApiCallFailed = "API call UpdateEndpoint failed with status code:%d, data:%s"
	errorUpdateEndpointDoesntExist = "API call UpdateEndpoint failed as endpoint with id:%d doesn't exist, data:%s"
)

func buildUpdateEndpointApiRequest(apiToken string, service string, endpoint Endpoint) (*http.Request, error) {
	jsonObject, err := buildUpdateEndpointRequest(endpoint)
	jsonBytes, err := json.Marshal(jsonObject)
	if err != nil {
		return nil, err
	}

	baseUrl := client.GetLogzioBaseUrl()
	id := endpoint.Id
	req, err := http.NewRequest(updateEndpointServiceMethod, fmt.Sprintf(updateEndpointServiceUrl, baseUrl, strings.ToLower(service), id), bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func buildUpdateEndpointRequest(endpoint Endpoint) (map[string]interface{}, error) {
	var updateEndpoint = map[string]interface{}{}

	updateEndpoint[fldEndpointTitle] = endpoint.Title
	updateEndpoint[fldEndpointDescription] = endpoint.Description

	if strings.EqualFold(EndpointTypeSlack, endpoint.EndpointType) {
		updateEndpoint[fldEndpointUrl] = endpoint.Url
	} else if strings.EqualFold(EndpointTypeCustom, endpoint.EndpointType) {
		updateEndpoint[fldEndpointUrl] = endpoint.Url
		updateEndpoint[fldEndpointMethod] = endpoint.Method
		headers := endpoint.Headers
		headerStrings := []string{}
		for k, v := range headers {
			headerStrings = append(headerStrings, fmt.Sprintf("%s=%s", k, v))
		}
		headerString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(headerStrings)), ","), "[]")
		updateEndpoint[fldEndpointHeaders] = headerString
		updateEndpoint[fldEndpointBodyTemplate] = endpoint.BodyTemplate
	} else if strings.EqualFold(endpoint.EndpointType, EndpointTypePagerDuty) {
		updateEndpoint[fldEndpointServiceKey] = endpoint.ServiceKey
	} else if strings.EqualFold(endpoint.EndpointType, EndpointTypeBigPanda) {
		updateEndpoint[fldEndpointApiToken] = endpoint.ApiToken
		updateEndpoint[fldEndpointAppKey] = endpoint.AppKey
	} else if strings.EqualFold(endpoint.EndpointType, EndpointTypeDataDog) {
		updateEndpoint[fldEndpointApiKey] = endpoint.ApiKey
	} else if strings.EqualFold(endpoint.EndpointType, EndpointTypeVictorOps) {
		updateEndpoint[fldEndpointRoutingKey] = endpoint.RoutingKey
		updateEndpoint[fldEndpointMessageType] = endpoint.MessageType
		updateEndpoint[fldEndpointServiceApiKey] = endpoint.ServiceApiKey
	} else {
		return nil, fmt.Errorf("don't recognise endpoint type %s", endpoint.EndpointType)
	}

	return updateEndpoint, nil
}

func (c *Endpoints) UpdateEndpoint(id int64, endpoint Endpoint) (*Endpoint, error) {

	endpoint.Id = id
	if jsonBytes, err, ok := c.makeEndpointRequest(endpoint, ValidateEndpointRequest, buildUpdateEndpointApiRequest, func(b []byte) error {
		if strings.Contains(fmt.Sprintf("%s", b), "Insufficient privileges") {
			return fmt.Errorf("API call %s failed for endpoint %d, data: %s", "UpdateEndpoint", id, b)
		}

		if strings.Contains(fmt.Sprintf("%s", b), "already exists") {
			return fmt.Errorf("API call %s failed for endpoint %d, data: %s", "UpdateEndpoint", id, b)
		}
		return nil
	}); ok {
		var target Endpoint
		err = json.Unmarshal(jsonBytes, &target)
		if err != nil {
			return nil, err
		}
		return &target, nil
	} else {
		return nil, err
	}
}
