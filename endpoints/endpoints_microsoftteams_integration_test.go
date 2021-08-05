package endpoints_test

import (
	"github.com/logzio/logzio_terraform_client/endpoints"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestIntegrationEndpoints_CreateEndpointMicrosoftTeams(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_microsoftteams"
		createEndpoint.Type = endpoints.EndpointTypeMicrosoftTeams
		createEndpoint.Url = testsUrl
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
			time.Sleep(time.Second * 1)
		}
	}
}

func TestIntegrationEndpoints_CreateEndpointMicrosoftTeamsEmptyUrl(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_microsoftteams_empy_url"
		createEndpoint.Type = endpoints.EndpointTypeMicrosoftTeams
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_CreateEndpointMicrosoftTeamsDuplicationError(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_microsoftteams_dup"
		createEndpoint.Type = endpoints.EndpointTypeMicrosoftTeams
		createEndpoint.Url = testsUrl
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
			time.Sleep(time.Second * 1)
			duplicate, err := underTest.CreateEndpoint(createEndpoint)
			assert.Error(t, err)
			assert.Nil(t, duplicate)
		}
	}
}

func TestIntegrationEndpoints_CreateEndpointMicrosoftTeamsNoTitle(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = ""
		createEndpoint.Type = endpoints.EndpointTypeMicrosoftTeams
		createEndpoint.Url = testsUrl
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_UpdateEndpointMicrosoftTeams(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_microsoftteams_to_update"
		createEndpoint.Type = endpoints.EndpointTypeMicrosoftTeams
		createEndpoint.Url = testsUrl
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
			time.Sleep(time.Second * 1)
			createEndpoint.Title = "updated_microsoftteams"
			createEndpoint.Description = "This is an UPDATED description"
			createEndpoint.Url = testsUrlUpdate
			updated, err := underTest.UpdateEndpoint(int64(endpoint.Id), createEndpoint)
			assert.NoError(t, err)
			assert.NotNil(t, updated)
			assert.Equal(t, endpoint.Id, updated.Id)
		}
	}
}

func TestIntegrationEndpoints_GetEndpointMicrosoftTeams(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_get_microsoftteams"
		createEndpoint.Type = endpoints.EndpointTypeMicrosoftTeams
		createEndpoint.Url = testsUrl
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
			time.Sleep(time.Second * 1)
			endpointFromGet, err := underTest.GetEndpoint(int64(endpoint.Id))
			assert.NoError(t, err)
			assert.NotNil(t, endpointFromGet)
			assert.Equal(t, endpoint.Id, endpointFromGet.Id)
			// TODO: change the following line if the type will be camelcase
			assert.Equal(t, "microsoft teams", strings.ToLower(endpointFromGet.Type))
			assert.Equal(t, createEndpoint.Title, endpointFromGet.Title)
			assert.Equal(t, createEndpoint.Url, endpointFromGet.Url)
		}
	}
}