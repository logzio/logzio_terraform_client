package endpoints

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpoints_CreateDeleteValidEndpoint(t *testing.T) {
	var endpoint *Endpoint
	var err error

	setupEndpointsTest()

	assert.NotNil(t, endpoints)

	if endpoints != nil {
		endpoint, err = endpoints.CreateEndpoint(deleteValidEndpoint())
		assert.Nil(t, err)

		err = endpoints.DeleteEndpoint(endpoint.Id)
		assert.NoError(t, err)

		_, err = endpoints.GetEndpoint(endpoint.Id)
		assert.Error(t, err)
	}

	teardownEndpointsTest()
}

func deleteValidEndpoint() Endpoint {
	return Endpoint{
		Title:        "deleteValidEndpoint",
		Description:  "my description",
		Url:          "https://this.is.com/some/webhook",
		EndpointType: "slack",
	}
}

// Tests create of an already existing endpoint (same titles)
func TestEndpointsCreateEndpointAlreadyExists(t *testing.T) {
	var endpoint *Endpoint
	var err error

	setupEndpointsTest()

	assert.NotNil(t, endpoints)

	if endpoints != nil {
		endpoint, err = endpoints.CreateEndpoint(createDuplicateEndpoint())
		assert.Nil(t, err)
		assert.NotNil(t, endpoint)
		createdEndpoints = append(createdEndpoints, endpoint.Id)

		endpoint, err = endpoints.CreateEndpoint(createDuplicateEndpoint())
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}

	teardownEndpointsTest()
}

func TestEndpointsCreateValidEndpoint(t *testing.T) {
	var endpoint *Endpoint
	var err error

	setupEndpointsTest()

	assert.NotNil(t, endpoints)

	if endpoints != nil {
		endpoint, err = endpoints.CreateEndpoint(createValidEndpoint())
		assert.Nil(t, err)
		createdEndpoints = append(createdEndpoints, endpoint.Id)

		selectedEndpoint, err := endpoints.GetEndpoint(endpoint.Id)
		assert.NoError(t, err)
		assert.NotNil(t, selectedEndpoint)
		assert.Equal(t, endpoint.Id, selectedEndpoint.Id)

		_, err = endpoints.GetEndpointByName(createValidEndpoint().Title)
		assert.NoError(t, err)

		_, err = endpoints.UpdateEndpoint(endpoint.Id, updateValidEndpoint())
		assert.NoError(t, err)

		updatedEndpoint, err := endpoints.GetEndpoint(endpoint.Id)
		assert.NoError(t, err)
		assert.Equal(t, endpoint.Id, updatedEndpoint.Id)
		assert.Equal(t, updateValidEndpoint().Title, updatedEndpoint.Title)
		assert.Equal(t, updateValidEndpoint().Url, updatedEndpoint.Url)
		assert.Equal(t, updateValidEndpoint().Description, updatedEndpoint.Description)
	}

	teardownEndpointsTest()
}

func TestEndpointsListEndpoints(t *testing.T) {
	assert.NotNil(t, endpoints)

	setupEndpointsTest()

	if endpoints != nil {
		endpoint, err := endpoints.CreateEndpoint(createValidEndpoint())
		list, err := endpoints.ListEndpoints()
		assert.NoError(t, err)
		assert.True(t, len(list) > 0)
		createdEndpoints = append(createdEndpoints, endpoint.Id)
	}

	teardownEndpointsTest()
}

func TestEndpointsCreateInvalidEndpoint(t *testing.T) {

	setupEndpointsTest()
	if assert.NotNil(t, endpoints) {
		invalidEndpoint := createInvalidEndpoint()
		endpoint, err := endpoints.CreateEndpoint(invalidEndpoint)
		assert.Nil(t, endpoint)
		assert.Error(t, err)
	}
}

func createDuplicateEndpoint() Endpoint {
	return Endpoint{
		Title:        "duplicateEndpoint",
		Description:  "my description",
		Url:          "https://this.is.com/some/webhook",
		EndpointType: "slack",
	}
}

func createValidEndpoint() Endpoint {
	return Endpoint{
		Title:        "validEndpoint",
		Description:  "my description",
		Url:          "https://this.is.com/some/webhook",
		EndpointType: "slack",
	}
}

func createInvalidEndpoint() Endpoint {
	return Endpoint{
		Title:        "invalidEndpoint",
		Description:  "my description",
		Url:          "https://someUrl",
		EndpointType: "slack",
	}
}

func updateValidEndpoint() Endpoint {
	return Endpoint{
		Title:        "updatedEndpoint",
		Description:  "my updated description",
		Url:          "https://this.is.com/some/other/webhook",
		EndpointType: "slack",
	}
}
