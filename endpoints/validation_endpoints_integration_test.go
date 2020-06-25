package endpoints_test

import (
	"testing"

	"github.com/logzio/logzio_terraform_client/endpoints"
	"github.com/stretchr/testify/assert"
)

func TestEndpoints_InvalidEndpoint(t *testing.T) {
	if endpoints.ValidateEndpointRequest(endpoints.Endpoint{
		Title: "title",
	}) {
		assert.Fail(t, "shouldn't be true")
	}
}

func TestEndpoints_ValidateSlackEndpoint(t *testing.T) {
	if !endpoints.ValidateEndpointRequest(endpoints.Endpoint{
		Title:        "title",
		Description:  "description",
		Url:          "url",
		EndpointType: endpoints.EndpointTypeSlack,
	}) {
		assert.Fail(t, "shouldn't be false")
	}
}

func TestEndpoints_ValidateCustomEndpoint(t *testing.T) {
	if !endpoints.ValidateEndpointRequest(endpoints.Endpoint{
		Title:        "title",
		Description:  "description",
		Url:          "url",
		Method:       "method",
		Headers:      map[string]string{"key": "value"},
		BodyTemplate: "{\"hello\":\"there\"}",
		EndpointType: endpoints.EndpointTypeCustom,
	}) {
		assert.Fail(t, "shouldn't be false")
	}
}

func TestEndpoints_ValidatePagerDuty(t *testing.T) {
	if !endpoints.ValidateEndpointRequest(endpoints.Endpoint{
		Title:        "title",
		Description:  "description",
		ServiceKey:   "serviceKey",
		EndpointType: endpoints.EndpointTypePagerDuty,
	}) {
		assert.Fail(t, "shouldn't be false")
	}
}

func TestEndpointsValidateBigPanda(t *testing.T) {
	if !endpoints.ValidateEndpointRequest(endpoints.Endpoint{
		Title:        "title",
		Description:  "description",
		ApiToken:     "ApiToken",
		AppKey:       "AppKey",
		EndpointType: endpoints.EndpointTypeBigPanda,
	}) {
		assert.Fail(t, "shouldn't be false")
	}
}

func TestEndpointsValidateDataDog(t *testing.T) {
	if !endpoints.ValidateEndpointRequest(endpoints.Endpoint{
		Title:        "title",
		Description:  "description",
		ApiKey:       "ApiKey",
		EndpointType: endpoints.EndpointTypeDataDog,
	}) {
		assert.Fail(t, "shouldn't be false")
	}
}

func TestEndpointsValidateVictorOps(t *testing.T) {
	if !endpoints.ValidateEndpointRequest(endpoints.Endpoint{
		Title:         "title",
		Description:   "description",
		RoutingKey:    "routingKey",
		MessageType:   "messageType",
		ServiceApiKey: "serviceApiKey",
		EndpointType:  endpoints.EndpointTypeVictorOps,
	}) {
		assert.Fail(t, "shouldn't be false")
	}
}
