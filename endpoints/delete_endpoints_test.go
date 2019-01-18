package endpoints

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpoints_CreateDeleteValidEndpoint(t *testing.T) {
	var endpoint *Endpoint
	var err error

	assert.NotNil(t, endpoints)

	if endpoints != nil {
		endpoint, err = endpoints.CreateEndpoint(deleteValidEndpoint())
		assert.Nil(t, err)

		err = endpoints.DeleteEndpoint(endpoint.Id)
		assert.NoError(t, err)

		_, err = endpoints.GetEndpoint(endpoint.Id)
		assert.Error(t, err)
	}
}

func deleteValidEndpoint() Endpoint {
	return Endpoint{
		Title:        "deleteValidEndpoint",
		Description:  "my description",
		Url:          "https://this.is.com/some/webhook",
		EndpointType: "slack",
	}
}
