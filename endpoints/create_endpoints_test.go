package endpoints

import (
	"github.com/jonboydell/logzio_client/test_utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var endpoints *Endpoints
var api_token string
var createdEndpoints []int64

func setup() {
	api_token = test_utils.GetApiToken()
	endpoints, _ = New(api_token)
	createdEndpoints = []int64{}
}

func shutdown() {
	for x := 0; x < len(createdEndpoints); x++ {
		endpoints.DeleteEndpoint(createdEndpoints[x])
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

// Tests create of an already existing endpoint (same titles)
func TestEndpoints_CreateEndpointAlreadyExists(t *testing.T) {
	var endpoint *Endpoint
	var err error

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
}

func TestEndpoints_CreateValidEndpoint(t *testing.T) {
	var endpoint *Endpoint
	var err error

	assert.NotNil(t, endpoints)

	if endpoints != nil {
		endpoint, err = endpoints.CreateEndpoint(createValidEndpoint())
		assert.Nil(t, err)
		createdEndpoints = append(createdEndpoints, endpoint.Id)

		selectedEndpoint, err := endpoints.GetEndpoint(endpoint.Id)
		assert.NoError(t, err)
		assert.NotNil(t, selectedEndpoint)
		assert.Equal(t, endpoint.Id, selectedEndpoint.Id)

		_, err = endpoints.updateEndpoint(endpoint.Id, updateValidEndpoint(), "slack")
		assert.NoError(t, err)

		updatedEndpoint, err := endpoints.GetEndpoint(endpoint.Id)
		assert.NoError(t, err)
		assert.Equal(t, endpoint.Id, updatedEndpoint.Id)
		assert.Equal(t, updateValidEndpoint().Title, updatedEndpoint.Title)
		assert.Equal(t, updateValidEndpoint().Url, updatedEndpoint.Url)
		assert.Equal(t, updateValidEndpoint().Description, updatedEndpoint.Description)
	}
}

func TestEndpoints_ListEndpoints(t *testing.T) {
	assert.NotNil(t, endpoints)

	if endpoints != nil {
		list, err := endpoints.ListEndpoints()
		assert.NoError(t, err)
		assert.True(t, len(list) > 0)
	}
}

func TestEndpoints_CreateInvalidEndpoint(t *testing.T) {
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
