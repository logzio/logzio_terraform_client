package endpoints

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpoints_buildCreateEndpointRequestSlack(t *testing.T) {
	result := buildCreateEndpointRequest(Endpoint{
		Title:        "title",
		Description:  "description",
		Url:          "url",
		EndpointType: endpointTypeSlack,
	})

	assert.Contains(t, result, fldEndpointTitle)
	assert.Contains(t, result, fldEndpointDescription)
	assert.Contains(t, result, fldEndpointUrl)
	assert.NotContains(t, result, fldEndpointType)
}

func TestEndpoints_buildCreateEndpointRequestCustom(t *testing.T) {
	result := buildCreateEndpointRequest(Endpoint{
		Title:        "title",
		Description:  "description",
		Url:          "url",
		Method:       "method",
		Headers:      map[string]string{"key": "value"},
		BodyTemplate: map[string]string{"key": "value"},
		EndpointType: endpointTypeCustom,
	})

	assert.Contains(t, result, fldEndpointTitle)
	assert.Contains(t, result, fldEndpointDescription)
	assert.Contains(t, result, fldEndpointUrl)
	assert.Contains(t, result, fldEndpointMethod)
	assert.Contains(t, result, fldEndpointHeaders)
	assert.Contains(t, result, fldEndpointBodyTemplate)
	assert.NotContains(t, result, fldEndpointType)
}

func TestEndpoints_buildCreateEndpointRequestPagerDuty(t *testing.T) {
	result := buildCreateEndpointRequest(Endpoint{
		Title:        "title",
		Description:  "description",
		ServiceKey:   "serviceKey",
		EndpointType: endpointTypePagerDuty,
	})
	assert.Contains(t, result, fldEndpointTitle)
	assert.Contains(t, result, fldEndpointDescription)
	assert.Contains(t, result, fldEndpointServiceKey)
	assert.NotContains(t, result, fldEndpointType)
}

func TestEndpoints_buildCreateEndpointRequestBigPanda(t *testing.T) {
	result := buildCreateEndpointRequest(Endpoint{
		Title:        "title",
		Description:  "description",
		ApiToken:     "ApiToken",
		AppKey:       "AppKey",
		EndpointType: endpointTypeBigPanda,
	})
	assert.Contains(t, result, fldEndpointTitle)
	assert.Contains(t, result, fldEndpointDescription)
	assert.Contains(t, result, fldEndpointApiToken)
	assert.Contains(t, result, fldEndpointAppKey)
	assert.NotContains(t, result, fldEndpointType)
}

func TestEndpoints_buildCreateEndpointRequestDataDog(t *testing.T) {
	result := buildCreateEndpointRequest(Endpoint{
		Title:        "title",
		Description:  "description",
		ApiKey:       "ApiKey",
		EndpointType: endpointTypeDataDog,
	})
	assert.Contains(t, result, fldEndpointTitle)
	assert.Contains(t, result, fldEndpointDescription)
	assert.Contains(t, result, fldEndpointApiKey)
	assert.NotContains(t, result, fldEndpointType)
}

func TestEndpoints_buildCreateEndpointRequestVictorOps(t *testing.T) {
	result := buildCreateEndpointRequest(Endpoint{
		Title:         "title",
		Description:   "description",
		RoutingKey:    "routingKey",
		MessageType:   "messageType",
		ServiceApiKey: "serviceApiKey",
		EndpointType:  endpointTypeVictorOps,
	})
	assert.Contains(t, result, fldEndpointTitle)
	assert.Contains(t, result, fldEndpointDescription)
	assert.Contains(t, result, fldEndpointRoutingKey)
	assert.Contains(t, result, fldEndpointMessageType)
	assert.Contains(t, result, fldEndpointServiceApiKey)
	assert.NotContains(t, result, fldEndpointType)
}
