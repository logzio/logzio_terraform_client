package endpoints

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpointsPagerDutyCreateUpdate(t *testing.T) {
	setupEndpointsTest()

	if assert.NotNil(t, endpoints) {
		endpoint, err := endpoints.CreateEndpoint(createPagerDutyEndpoint())
		assert.NotNil(t, endpoint)

		if assert.NoError(t, err) {
			createdEndpoints = append(createdEndpoints, endpoint.Id)
			endpoint, err = endpoints.UpdateEndpoint(endpoint.Id, createPagerDutyEndpoint())
			assert.NotNil(t, endpoint)
			assert.NoError(t, err)
		}
	}

	teardownEndpointsTest()
}

func createPagerDutyEndpoint() Endpoint {
	return Endpoint{
		Title:        "validEndpoint",
		Description:  "my description",
		EndpointType: "pager-duty",
		ServiceKey:   "my_service_key",
	}
}
