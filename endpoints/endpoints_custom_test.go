package endpoints_test

import (
	"github.com/jonboydell/logzio_client/endpoints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpointsCustomCreateUpdate(t *testing.T) {
	underTest, err := setupEndpointsTest()
	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(createCustomEndpoint())
		if assert.NotNil(t, endpoint) {
			assert.NoError(t, err)
			endpoint, err = underTest.UpdateEndpoint(endpoint.Id, createUpdatedCustomEndpoint())
			assert.NotNil(t, endpoint)
			assert.NoError(t, err)

			underTest.DeleteEndpoint(endpoint.Id)
		}
	}
}

func TestEndpointsCustomCreateDuplicate(t *testing.T) {
	underTest, err := setupEndpointsTest()
	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(createCustomEndpoint())
		if assert.NotNil(t, endpoint) {
			assert.NoError(t, err)
			endpoint, err = underTest.CreateEndpoint(createCustomEndpoint())
			assert.Nil(t, endpoint)
			assert.Error(t, err)
		}
	}
}

func createCustomEndpoint() endpoints.Endpoint {
	return endpoints.Endpoint{
		Title:        "customEndpoint",
		Method:       "POST",
		Description:  "my description",
		Url:          "https://this.is.com/some/other/webhook",
		EndpointType: "custom",
		Headers:      map[string]string{"hello": "there", "header": "two"},
		BodyTemplate: map[string]string{"hello": "there", "header": "two"},
	}
}

func createUpdatedCustomEndpoint() endpoints.Endpoint {
	return endpoints.Endpoint{
		Title:        "customEndpoint",
		Method:       "POST",
		Description:  "some updated description",
		Url:          "https://this.is.com/some/other/webhook",
		EndpointType: "custom",
		Headers:      map[string]string{"hello": "there", "header": "two"},
		BodyTemplate: map[string]string{"hello": "there", "header": "two"},
	}
}
