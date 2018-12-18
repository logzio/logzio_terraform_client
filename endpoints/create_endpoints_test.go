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
func TestCreateEndpointAlreadyExists(t *testing.T) {
	var endpoint *EndpointType
	var err error

	assert.NotNil(t, endpoints)

	if endpoints != nil {
		endpoint, err = endpoints.createEndpoint(createDuplicateEndpoint(), "slack")
		assert.Nil(t, err)
		assert.NotNil(t, endpoint)
		createdEndpoints = append(createdEndpoints, endpoint.Id)

		endpoint, err = endpoints.createEndpoint(createDuplicateEndpoint(), "slack")
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestCreateValidEndpoint(t *testing.T) {
	var endpoint *EndpointType
	var err error

	assert.NotNil(t, endpoints)

	if endpoints != nil {
		endpoint, err = endpoints.createEndpoint(createValidEndpoint(), "slack")
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

func TestCreateInvalidEndpoint(t *testing.T) {
	if assert.NotNil(t, endpoints) {
		invalidEndpoint := createInvalidEndpoint()
		endpoint, err := endpoints.createEndpoint(invalidEndpoint, "slack")
		assert.Nil(t, endpoint)
		assert.Error(t, err)
	}
}

func createDuplicateEndpoint() EndpointType {
	return EndpointType{
		Title:       "duplicateEndpoint",
		Description: "my description",
		Url:         "https://this.is.com/some/webhook",
	}
}

func createValidEndpoint() EndpointType {
	return EndpointType{
		Title:       "validEndpoint",
		Description: "my description",
		Url:         "https://this.is.com/some/webhook",
	}
}

func createInvalidEndpoint() EndpointType {
	return EndpointType{
		Title:       "invalidEndpoint",
		Description: "my description",
		Url:         "https://someUrl",
	}
}

func updateValidEndpoint() EndpointType {
	return EndpointType{
		Title:       "updatedEndpoint",
		Description: "my updated description",
		Url:         "https://this.is.com/some/other/webhook",
	}
}
