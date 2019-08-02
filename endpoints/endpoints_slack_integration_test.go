// +build integration

package endpoints_test

import (
	"github.com/jonboydell/logzio_client/endpoints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationEndpoints_CreateDeleteGetValidEndpoint(t *testing.T) {
	var endpoint *endpoints.Endpoint
	var err error

	underTest, err := setupEndpointsIntegrationTest()

	if assert.NoError(t, err) {
		endpoint, err = underTest.CreateEndpoint(endpoints.Endpoint{
			Title:        "slackcreatedeletevalidendpoint",
			Description:  "my description",
			Url:          "https://this.is.com/some/webhook",
			EndpointType: "slack",
		})
		assert.Nil(t, err)

		err = underTest.DeleteEndpoint(endpoint.Id)
		assert.NoError(t, err)

		_, err = underTest.GetEndpoint(endpoint.Id)
		assert.Error(t, err)
	}
}

// Tests create of an already existing endpoint (same titles)
func TestIntegrationEndpoints_CreateDuplicateEndpoint(t *testing.T) {
	var endpoint *endpoints.Endpoint
	var err error

	underTest, err := setupEndpointsIntegrationTest()

	if assert.NoError(t, err) {
		endpoint, err = underTest.CreateEndpoint(endpoints.Endpoint{
			Title:        "slackcreateduplicateendpoint",
			Description:  "my description",
			Url:          "https://this.is.com/some/webhook",
			EndpointType: "slack",
		})
		if assert.NoError(t, err) {
			duplicate, err := underTest.CreateEndpoint(endpoints.Endpoint{
				Title:        "slackcreateduplicateendpoint",
				Description:  "my description",
				Url:          "https://this.is.com/some/webhook",
				EndpointType: "slack",
			})
			assert.Error(t, err)
			assert.Nil(t, duplicate)
		}
		err = underTest.DeleteEndpoint(endpoint.Id)
		assert.NoError(t, err)
	}
}

func TestIntegrationEndpoints_ListEndpoints(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(endpoints.Endpoint{
			Title:        "slacklistendpoints",
			Description:  "my description",
			Url:          "https://this.is.com/some/webhook",
			EndpointType: "slack",
		})
		list, err := underTest.ListEndpoints()
		assert.NoError(t, err)
		assert.True(t, len(list) > 0)
		err = underTest.DeleteEndpoint(endpoint.Id)
		assert.NoError(t, err)
	}
}

func TestIntegrationEndpoints_CreateInvalidEndpoint(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(endpoints.Endpoint{
			Title:        "slackinvalidEndpoint",
			Description:  "my description",
			Url:          "https://someUrl",
			EndpointType: "slack",
		})
		assert.Nil(t, endpoint)
		assert.Error(t, err)
	}
}

func TestIntegrationEndpoints_UpdateEndpoint(t *testing.T) {
	var endpoint *endpoints.Endpoint
	var err error

	underTest, err := setupEndpointsIntegrationTest()

	if assert.NoError(t, err) && assert.NotNil(t, underTest) {
		endpoint, err = underTest.CreateEndpoint(endpoints.Endpoint{
			Title:        "slackupdatedendpoint",
			Description:  "my description",
			Url:          "https://this.is.com/some/webhook",
			EndpointType: "slack",
		})
		assert.NoError(t, err)
		assert.NotNil(t, endpoint)

		updatedEndpoint, err := underTest.UpdateEndpoint(endpoint.Id, endpoints.Endpoint{
			Title:        "slackupdatedupdatedendpoint",
			Description:  "my updated description",
			Url:          "https://this.is.com/some/other/webhook",
			EndpointType: "slack",
		})
		assert.NoError(t, err)
		assert.NotNil(t, updatedEndpoint)

		readEndpoint, err := underTest.GetEndpoint(updatedEndpoint.Id)
		assert.NoError(t, err)
		assert.NotNil(t, readEndpoint)

		assert.Equal(t, "slackupdatedupdatedendpoint", readEndpoint.Title)

		err = underTest.DeleteEndpoint(endpoint.Id)
		assert.NoError(t, err)
	}
}
