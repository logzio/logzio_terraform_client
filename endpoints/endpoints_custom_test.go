package endpoints

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpointsCustomCreateUpdate(t *testing.T) {
	setupEndpointsTest()
	if assert.NotNil(t, endpoints) {
		endpoint, err := endpoints.CreateEndpoint(createCustomEndpoint())
		if assert.NotNil(t, endpoint) {
			assert.NoError(t, err)
			createdEndpoints = append(createdEndpoints, endpoint.Id)
			endpoint, err = endpoints.UpdateEndpoint(endpoint.Id, createUpdatedCustomEndpoint())
			assert.NotNil(t, endpoint)
			assert.NoError(t, err)
		}
	}
	teardownEndpointsTest()
}

func TestEndpointsCustomCreateDuplicate(t *testing.T) {
	setupEndpointsTest()
	if assert.NotNil(t, endpoints) {
		endpoint, err := endpoints.CreateEndpoint(createCustomEndpoint())
		if assert.NotNil(t, endpoint) {
			assert.NoError(t, err)
			createdEndpoints = append(createdEndpoints, endpoint.Id)
			endpoint, err = endpoints.CreateEndpoint(createCustomEndpoint())
			assert.Nil(t, endpoint)
			assert.Error(t, err)
		}
	}
	teardownEndpointsTest()
}

func createCustomEndpoint() Endpoint {
	return Endpoint{
		Title:        "customEndpoint",
		Method:       "POST",
		Description:  "my description",
		Url:          "https://this.is.com/some/other/webhook",
		EndpointType: "custom",
		Headers:      map[string]string{"hello": "there", "header": "two"},
		BodyTemplate: map[string]string{"hello": "there", "header": "two"},
	}
}

func createUpdatedCustomEndpoint() Endpoint {
	return Endpoint{
		Title:        "customEndpoint",
		Method:       "POST",
		Description:  "some updated description",
		Url:          "https://this.is.com/some/other/webhook",
		EndpointType: "custom",
		Headers:      map[string]string{"hello": "there", "header": "two"},
		BodyTemplate: map[string]string{"hello": "there", "header": "two"},
	}
}
