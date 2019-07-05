package endpoints_test

import (
	"github.com/jonboydell/logzio_client/endpoints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpointsInvalidEndpoint(t *testing.T) {
	if err, ok := endpoints.ValidateEndpointRequest(endpoints.Endpoint{
		Title: "title",
	}); ok {
		assert.Fail(t, "shouldn't be true")
	} else {
		assert.Error(t, err)
	}
}

func TestEndpointsValidateSlackEndpoint(t *testing.T) {
	if err, ok := endpoints.ValidateEndpointRequest(endpoints.Endpoint{
		Title:        "title",
		Description:  "description",
		Url:          "url",
		EndpointType: endpoints.EndpointTypeSlack,
	}); ok {
		assert.NoError(t, err)
	} else {
		assert.Fail(t, "shouldn't be false")
	}
}

func TestEndpointsValidateCustomEndpoint(t *testing.T) {
	if err, ok := endpoints.ValidateEndpointRequest(endpoints.Endpoint{
		Title:        "title",
		Description:  "description",
		Url:          "url",
		Method:       "method",
		Headers:      map[string]string{"key": "value"},
		BodyTemplate: "{\"hello\":\"there\"}",
		EndpointType: endpoints.EndpointTypeCustom,
	}); ok {
		assert.NoError(t, err)
	} else {
		assert.Fail(t, "shouldn't be false")
	}
}

func TestEndpointsValidatePagerDuty(t *testing.T) {
	if err, ok := endpoints.ValidateEndpointRequest(endpoints.Endpoint{
		Title:        "title",
		Description:  "description",
		ServiceKey:   "serviceKey",
		EndpointType: endpoints.EndpointTypePagerDuty,
	}); ok {
		assert.NoError(t, err)
	} else {
		assert.Fail(t, "shouldn't be false")
	}
}

func TestEndpointsValidateBigPanda(t *testing.T) {
	if err, ok := endpoints.ValidateEndpointRequest(endpoints.Endpoint{
		Title:        "title",
		Description:  "description",
		ApiToken:     "ApiToken",
		AppKey:       "AppKey",
		EndpointType: endpoints.EndpointTypeBigPanda,
	}); ok {
		assert.NoError(t, err)
	} else {
		assert.Fail(t, "shouldn't be false")
	}
}

func TestEndpointsValidateDataDog(t *testing.T) {
	if err, ok := endpoints.ValidateEndpointRequest(endpoints.Endpoint{
		Title:        "title",
		Description:  "description",
		ApiKey:       "ApiKey",
		EndpointType: endpoints.EndpointTypeDataDog,
	}); ok {
		assert.NoError(t, err)
	} else {
		assert.Fail(t, "shouldn't be false")
	}
}

func TestEndpointsValidateVictorOps(t *testing.T) {
	if err, ok := endpoints.ValidateEndpointRequest(endpoints.Endpoint{
		Title:         "title",
		Description:   "description",
		RoutingKey:    "routingKey",
		MessageType:   "messageType",
		ServiceApiKey: "serviceApiKey",
		EndpointType:  endpoints.EndpointTypeVictorOps,
	}); ok {
		assert.NoError(t, err)
	} else {
		assert.Fail(t, "shouldn't be false")
	}
}
