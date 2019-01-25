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
		EndpointType: EndpointTypeSlack,
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
		BodyTemplate: "{ \"hello\" : \"there\"}",
		EndpointType: EndpointTypeCustom,
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
		EndpointType: EndpointTypePagerDuty,
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
		EndpointType: EndpointTypeBigPanda,
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
		EndpointType: EndpointTypeDataDog,
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
		EndpointType:  EndpointTypeVictorOps,
	})
	assert.Contains(t, result, fldEndpointTitle)
	assert.Contains(t, result, fldEndpointDescription)
	assert.Contains(t, result, fldEndpointRoutingKey)
	assert.Contains(t, result, fldEndpointMessageType)
	assert.Contains(t, result, fldEndpointServiceApiKey)
	assert.NotContains(t, result, fldEndpointType)
}
