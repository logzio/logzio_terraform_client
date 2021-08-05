package endpoints_test

import (
	"github.com/logzio/logzio_terraform_client/endpoints"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestIntegrationEndpoints_CreateEndpointServiceNow(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_servicenow"
		createEndpoint.Type = endpoints.EndpointTypeServiceNow
		createEndpoint.Username = "someUsername"
		createEndpoint.Password = "somePassword"
		createEndpoint.Url = testsUrl
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
			time.Sleep(time.Second * 1)
		}
	}
}

func TestIntegrationEndpoints_CreateEndpointServiceNowEmptyUsername(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_servicenow_empty_username"
		createEndpoint.Type = endpoints.EndpointTypeServiceNow
		createEndpoint.Password = "somePassword"
		createEndpoint.Url = testsUrl
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_CreateEndpointServiceNowEmptyPassword(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_servicenow_empty_password"
		createEndpoint.Type = endpoints.EndpointTypeServiceNow
		createEndpoint.Username = "someUsername"
		createEndpoint.Url = testsUrl
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_CreateEndpointServiceNowEmptyUrl(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_servicenow_empty_url"
		createEndpoint.Type = endpoints.EndpointTypeServiceNow
		createEndpoint.Username = "someUsername"
		createEndpoint.Password = "somePassword"
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_CreateEndpointServiceNowDuplicationError(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_servicenow_dup"
		createEndpoint.Type = endpoints.EndpointTypeServiceNow
		createEndpoint.Username = "someUsername"
		createEndpoint.Password = "somePassword"
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

func TestIntegrationEndpoints_CreateEndpointServiceNowNoTitle(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = ""
		createEndpoint.Type = endpoints.EndpointTypeServiceNow
		createEndpoint.Username = "someUsername"
		createEndpoint.Password = "somePassword"
		createEndpoint.Url = testsUrl
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_UpdateEndpointServiceNow(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_servicenow_to_update"
		createEndpoint.Type = endpoints.EndpointTypeServiceNow
		createEndpoint.Username = "someUsername"
		createEndpoint.Password = "somePassword"
		createEndpoint.Url = testsUrl
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
			time.Sleep(time.Second * 1)
			createEndpoint.Title = "updated_servicenow"
			createEndpoint.Description = "This is an UPDATED description"
			createEndpoint.Username = "updatedUsername"
			createEndpoint.Password = "updatedPassword"
			createEndpoint.Url = testsUrlUpdate
			updated, err := underTest.UpdateEndpoint(int64(endpoint.Id), createEndpoint)
			assert.NoError(t, err)
			assert.NotNil(t, updated)
			assert.Equal(t, endpoint.Id, updated.Id)
		}
	}
}

func TestIntegrationEndpoints_GetEndpointServiceNow(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_get_servicenow"
		createEndpoint.Type = endpoints.EndpointTypeServiceNow
		createEndpoint.Username = "someUsername"
		createEndpoint.Password = "somePassword"
		createEndpoint.Url = testsUrl
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
			time.Sleep(time.Second * 1)
			endpointFromGet, err := underTest.GetEndpoint(int64(endpoint.Id))
			assert.NoError(t, err)
			assert.NotNil(t, endpointFromGet)
			assert.Equal(t, endpoint.Id, endpointFromGet.Id)
			assert.Equal(t, createEndpoint.Type, strings.ToLower(endpointFromGet.Type))
			assert.Equal(t, createEndpoint.Title, endpointFromGet.Title)
			assert.Equal(t, createEndpoint.Username, endpointFromGet.Username)
			assert.Equal(t, createEndpoint.Password, endpointFromGet.Password)
			assert.Equal(t, createEndpoint.Url, endpointFromGet.Url)
		}
	}
}