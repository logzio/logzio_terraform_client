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

const updateEndpointServiceUrl string = "%s/v1/endpoints/%s/%d"
const updateEndpointServiceMethod string = http.MethodPut
const updateEndpointMethodSuccess int = 200

const errorUpdateEndpointApiCallFailed = "API call UpdateEndpoint failed with status code:%d, data:%s"
const errorUpdateEndpointDoesntExist = "API call UpdateEndpoint failed as endpoint with id:%d doesn't exist, data:%s"

func buildUpdateEndpointApiRequest(apiToken string, id int64, service string, jsonObject map[string]interface{}) (*http.Request, error) {
	jsonBytes, err := json.Marshal(jsonObject)
	if err != nil {
		return nil, err
	}

	baseUrl := client.GetLogzioBaseUrl()
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
		updateEndpoint[fldEndpointHeaders] = endpoint.Headers
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
	err := ValidateEndpointRequest(endpoint)
	if err != nil {
		return nil, err
	}

	updateEndpoint, err := buildUpdateEndpointRequest(endpoint)
	if err != nil {
		return nil, err
	}

	req, _ := buildUpdateEndpointApiRequest(c.ApiToken, id, endpoint.EndpointType, updateEndpoint)

	var httpClient http.Client
	resp, _ := httpClient.Do(req)
	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{updateEndpointMethodSuccess}) {
		return nil, fmt.Errorf(errorUpdateEndpointApiCallFailed, resp.StatusCode, jsonBytes)
	}

	var target Endpoint
	json.Unmarshal(jsonBytes, &target)

	if len(target.Message) > 0 {
		return nil, fmt.Errorf(errorUpdateEndpointDoesntExist, resp.StatusCode, target.Message)
	}

	return &target, nil
}
