// +build integration

package endpoints_test

import (
	"github.com/jonboydell/logzio_client/endpoints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpointsVictorOpsCreateUpdate(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(createVictorOpsEndpoint())
		assert.NotNil(t, endpoint)
		assert.NoError(t, err)

		endpoint, err = underTest.UpdateEndpoint(endpoint.Id, updateVictorOpsEndpoint())
		assert.NotNil(t, endpoint)
		assert.NoError(t, err)

		underTest.DeleteEndpoint(endpoint.Id)
	}
}

func createVictorOpsEndpoint() endpoints.Endpoint {
	return endpoints.Endpoint{
		Title:         "vopsendpoint",
		Description:   "description",
		EndpointType:  "victorops",
		RoutingKey:    "routing_key",
		MessageType:   "message_type",
		ServiceApiKey: "service_api_key",
	}
}

func updateVictorOpsEndpoint() endpoints.Endpoint {
	return endpoints.Endpoint{
		Title:         "vopsupdatedendpoint",
		Description:   "updated description",
		EndpointType:  "victorops",
		RoutingKey:    "updated_routing_key",
		MessageType:   "updated_message_type",
		ServiceApiKey: "updated_service_api_key",
	}
}
