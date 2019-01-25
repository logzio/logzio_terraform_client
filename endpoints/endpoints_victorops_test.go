package endpoints

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpointsVictorOpsCreateUpdate(t *testing.T) {
	setupEndpointsTest()

	if assert.NotNil(t, endpoints) {
		endpoint, err := endpoints.CreateEndpoint(createVictorOpsEndpoint())
		assert.NotNil(t, endpoint)
		assert.NoError(t, err)
		createdEndpoints = append(createdEndpoints, endpoint.Id)

		endpoint, err = endpoints.UpdateEndpoint(endpoint.Id, updateVictorOpsEndpoint())
		assert.NotNil(t, endpoint)
		assert.NoError(t, err)
	}

	teardownEndpointsTest()
}

func createVictorOpsEndpoint() Endpoint {
	return Endpoint{
		Title:         "vops endpoint",
		Description:   "description",
		EndpointType:  "victorops",
		RoutingKey:    "routing_key",
		MessageType:   "message_type",
		ServiceApiKey: "service_api_key",
	}
}

func updateVictorOpsEndpoint() Endpoint {
	return Endpoint{
		Title:         "vops updated endpoint",
		Description:   "updated description",
		EndpointType:  "victorops",
		RoutingKey:    "updated_routing_key",
		MessageType:   "updated_message_type",
		ServiceApiKey: "updated_service_api_key",
	}
}
