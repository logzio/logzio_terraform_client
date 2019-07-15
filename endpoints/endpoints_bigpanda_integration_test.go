// +build integration

package endpoints_test

import (
	"github.com/jonboydell/logzio_client/endpoints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpointsBigPandaCreateUpdate(t *testing.T) {
	underTest, err := setupEndpointsTest()
	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(createBigPandaEndpoint())

		if assert.NoError(t, err) {
			updatedEndpoint, err := underTest.UpdateEndpoint(endpoint.Id, updateBigPandaEndpoint())
			assert.NotNil(t, updatedEndpoint)
			assert.NoError(t, err)

			err = underTest.DeleteEndpoint(endpoint.Id)
			assert.NoError(t, err)
		}
	}
}

func createBigPandaEndpoint() endpoints.Endpoint {
	return endpoints.Endpoint{
		Title:        "bigpandaendpoint",
		Description:  "description",
		EndpointType: "big-panda",
		ApiToken:     "api_token",
		AppKey:       "app_key",
	}
}

func updateBigPandaEndpoint() endpoints.Endpoint {
	return endpoints.Endpoint{
		Title:        "bigpandaupdatedendpoint",
		Description:  "updated description",
		EndpointType: "big-panda",
		ApiToken:     "updated_api_token",
		AppKey:       "updated_app_key",
	}
}
