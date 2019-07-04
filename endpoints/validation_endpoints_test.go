package endpoints

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpointsInvalidEndpoint(t *testing.T) {
	if err, ok := ValidateEndpointRequest(Endpoint{
		Title: "title",
	}); ok {
		assert.Fail(t, "shouldn't be true")
	} else {
		assert.Error(t, err)
	}
}

func TestEndpointsValidateSlackEndpoint(t *testing.T) {
	if err, ok := ValidateEndpointRequest(Endpoint{
		Title:        "title",
		Description:  "description",
		Url:          "url",
		EndpointType: EndpointTypeSlack,
	}); ok {
		assert.NoError(t, err)
	} else {
		assert.Fail(t, "shouldn't be false")
	}
}

func TestEndpointsValidateCustomEndpoint(t *testing.T) {
	if err, ok := ValidateEndpointRequest(Endpoint{
		Title:        "title",
		Description:  "description",
		Url:          "url",
		Method:       "method",
		Headers:      map[string]string{"key": "value"},
		BodyTemplate: "{\"hello\":\"there\"}",
		EndpointType: EndpointTypeCustom,
	}); ok {
		assert.NoError(t, err)
	} else {
		assert.Fail(t, "shouldn't be false")
	}
}

func TestEndpointsValidatePagerDuty(t *testing.T) {
	if err, ok := ValidateEndpointRequest(Endpoint{
		Title:        "title",
		Description:  "description",
		ServiceKey:   "serviceKey",
		EndpointType: EndpointTypePagerDuty,
	}); ok {
		assert.NoError(t, err)
	} else {
		assert.Fail(t, "shouldn't be false")
	}
}

func TestEndpointsValidateBigPanda(t *testing.T) {
	if err, ok := ValidateEndpointRequest(Endpoint{
		Title:        "title",
		Description:  "description",
		ApiToken:     "ApiToken",
		AppKey:       "AppKey",
		EndpointType: EndpointTypeBigPanda,
	}); ok {
		assert.NoError(t, err)
	} else {
		assert.Fail(t, "shouldn't be false")
	}
}

func TestEndpointsValidateDataDog(t *testing.T) {
	if err, ok := ValidateEndpointRequest(Endpoint{
		Title:        "title",
		Description:  "description",
		ApiKey:       "ApiKey",
		EndpointType: EndpointTypeDataDog,
	}); ok {
		assert.NoError(t, err)
	} else {
		assert.Fail(t, "shouldn't be false")
	}
}

func TestEndpointsValidateVictorOps(t *testing.T) {
	if err, ok := ValidateEndpointRequest(Endpoint{
		Title:         "title",
		Description:   "description",
		RoutingKey:    "routingKey",
		MessageType:   "messageType",
		ServiceApiKey: "serviceApiKey",
		EndpointType:  EndpointTypeVictorOps,
	}); ok {
		assert.NoError(t, err)
	} else {
		assert.Fail(t, "shouldn't be false")
	}
}
