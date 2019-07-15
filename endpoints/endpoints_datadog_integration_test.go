// +build integration

package endpoints_test

import (
	"github.com/jonboydell/logzio_client/endpoints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpointsDataDogCreateUpdate(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(createDataDogEndpoint())
		assert.NotNil(t, endpoint)
		if assert.NoError(t, err) {
			endpoint, err = underTest.UpdateEndpoint(endpoint.Id, updateDataDogEndpoint())
			assert.NotNil(t, endpoint)
			assert.NoError(t, err)
		}
		underTest.DeleteEndpoint(endpoint.Id)
	}
}

func createDataDogEndpoint() endpoints.Endpoint {
	return endpoints.Endpoint{
		Title:        "datadogendpoint",
		Description:  "description",
		EndpointType: "data-dog",
		ApiKey:       "api_key",
	}
}

func updateDataDogEndpoint() endpoints.Endpoint {
	return endpoints.Endpoint{
		Title:        "datadogupdatedendpoint",
		Description:  "updated description",
		EndpointType: "data-dog",
		ApiKey:       "updated_api_key",
	}
}
