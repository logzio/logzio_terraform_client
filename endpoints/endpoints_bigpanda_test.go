package endpoints

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpointsBigPandaCreateUpdate(t *testing.T) {
	setupEndpointsTest()
	if assert.NotNil(t, endpoints) {
		endpoint, err := endpoints.CreateEndpoint(createBigPandaEndpoint())

		if assert.NotNil(t, endpoint) {
			assert.NoError(t, err)
			createdEndpoints = append(createdEndpoints, endpoint.Id)
			endpoint, err = endpoints.UpdateEndpoint(endpoint.Id, updateBigPandaEndpoint())
			assert.NotNil(t, endpoint)
			assert.NoError(t, err)
		}
	}
	teardownEndpointsTest()
}

func createBigPandaEndpoint() Endpoint {
	return Endpoint{
		Title:        "endpoint",
		Description:  "description",
		EndpointType: "big-panda",
		ApiToken:     "api_token",
		AppKey:       "app_key",
	}
}

func updateBigPandaEndpoint() Endpoint {
	return Endpoint{
		Title:        "updated endpoint",
		Description:  "updated description",
		EndpointType: "big-panda",
		ApiToken:     "updated_api_token",
		AppKey:       "updated_app_key",
	}
}
