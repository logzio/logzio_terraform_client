package endpoints_test

import (
	"github.com/jonboydell/logzio_client/endpoints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpoints_CreateDeleteValidEndpoint(t *testing.T) {
	var endpoint *endpoints.Endpoint
	var err error

	underTest, err := setupEndpointsTest()

	if assert.NoError(t, err) {
		endpoint, err = underTest.CreateEndpoint(deleteValidEndpoint())
		assert.Nil(t, err)

		err = underTest.DeleteEndpoint(endpoint.Id)
		assert.NoError(t, err)

		_, err = underTest.GetEndpoint(endpoint.Id)
		assert.Error(t, err)
	}
}

func deleteValidEndpoint() endpoints.Endpoint {
	return endpoints.Endpoint{
		Title:        "deleteValidEndpoint",
		Description:  "my description",
		Url:          "https://this.is.com/some/webhook",
		EndpointType: "slack",
	}
}

// Tests create of an already existing endpoint (same titles)
func TestEndpointsCreateEndpointAlreadyExists(t *testing.T) {
	var endpoint *endpoints.Endpoint
	var err error

	underTest, err := setupEndpointsTest()

	if assert.NoError(t, err) {
		endpoint, err = underTest.CreateEndpoint(createDuplicateEndpoint())
		assert.Nil(t, err)
		if assert.NotNil(t, endpoint) {
			duplicate, err := underTest.CreateEndpoint(createDuplicateEndpoint())
			assert.Error(t, err)
			assert.Nil(t, duplicate)
		}
		underTest.DeleteEndpoint(endpoint.Id)
	}
}

func TestEndpointsCreateValidEndpoint(t *testing.T) {
	var endpoint *endpoints.Endpoint
	var err error

	underTest, err := setupEndpointsTest()

	if assert.NoError(t, err) {
		endpoint, err = underTest.CreateEndpoint(createValidEndpoint())
		assert.Nil(t, err)

		selectedEndpoint, err := underTest.GetEndpoint(endpoint.Id)
		assert.NoError(t, err)
		assert.NotNil(t, selectedEndpoint)
		assert.Equal(t, endpoint.Id, selectedEndpoint.Id)

		_, err = underTest.GetEndpointByName(createValidEndpoint().Title)
		assert.NoError(t, err)

		_, err = underTest.UpdateEndpoint(endpoint.Id, updateValidEndpoint())
		assert.NoError(t, err)

		updatedEndpoint, err := underTest.GetEndpoint(endpoint.Id)
		assert.NoError(t, err)
		assert.Equal(t, endpoint.Id, updatedEndpoint.Id)
		assert.Equal(t, updateValidEndpoint().Title, updatedEndpoint.Title)
		assert.Equal(t, updateValidEndpoint().Url, updatedEndpoint.Url)
		assert.Equal(t, updateValidEndpoint().Description, updatedEndpoint.Description)
	}
}

func TestEndpointsListEndpoints(t *testing.T) {
	underTest, err := setupEndpointsTest()
	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(createValidEndpoint())
		list, err := underTest.ListEndpoints()
		assert.NoError(t, err)
		assert.True(t, len(list) > 0)
		underTest.DeleteEndpoint(endpoint.Id)
	}
}

func TestEndpointsCreateInvalidEndpoint(t *testing.T) {
	underTest, err := setupEndpointsTest()
	if assert.NoError(t, err) {
		invalidEndpoint := createInvalidEndpoint()
		endpoint, err := underTest.CreateEndpoint(invalidEndpoint)
		assert.Nil(t, endpoint)
		assert.Error(t, err)
	}
}

func createDuplicateEndpoint() endpoints.Endpoint {
	return endpoints.Endpoint{
		Title:        "duplicateEndpoint",
		Description:  "my description",
		Url:          "https://this.is.com/some/webhook",
		EndpointType: "slack",
	}
}

func createValidEndpoint() endpoints.Endpoint {
	return endpoints.Endpoint{
		Title:        "validEndpoint",
		Description:  "my description",
		Url:          "https://this.is.com/some/webhook",
		EndpointType: "slack",
	}
}

func createInvalidEndpoint() endpoints.Endpoint {
	return endpoints.Endpoint{
		Title:        "invalidEndpoint",
		Description:  "my description",
		Url:          "https://someUrl",
		EndpointType: "slack",
	}
}

func updateValidEndpoint() endpoints.Endpoint {
	return endpoints.Endpoint{
		Title:        "updatedEndpoint",
		Description:  "my updated description",
		Url:          "https://this.is.com/some/other/webhook",
		EndpointType: "slack",
	}
}
