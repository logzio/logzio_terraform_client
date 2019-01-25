package endpoints

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpointsDataDogCreateUpdate(t *testing.T) {
	setupEndpointsTest()

	if assert.NotNil(t, endpoints) {
		endpoint, err := endpoints.CreateEndpoint(createDataDogEndpoint())
		assert.NotNil(t, endpoint)

		createdEndpoints = append(createdEndpoints, endpoint.Id)
		if assert.NoError(t, err) {
			endpoint, err = endpoints.UpdateEndpoint(endpoint.Id, updateDataDogEndpoint())
			assert.NotNil(t, endpoint)
			assert.NoError(t, err)
		}
	}

	teardownEndpointsTest()
}

func createDataDogEndpoint() Endpoint {
	return Endpoint{
		Title:        "endpoint",
		Description:  "description",
		EndpointType: "data-dog",
		ApiKey:       "api_key",
	}
}

func updateDataDogEndpoint() Endpoint {
	return Endpoint{
		Title:        "updated endpoint",
		Description:  "updated description",
		EndpointType: "data-dog",
		ApiKey:       "updated_api_key",
	}
}
