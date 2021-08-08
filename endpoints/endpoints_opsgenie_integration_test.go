package endpoints_test

import (
	"github.com/logzio/logzio_terraform_client/endpoints"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestIntegrationEndpoints_CreateEndpointOpsGenie(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_opsgenie"
		createEndpoint.Type = endpoints.EndpointTypeOpsGenie
		createEndpoint.ApiKey = "someApiKey"
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
			time.Sleep(time.Second * 1)
		}
	}
}

func TestIntegrationEndpoints_CreateEndpointOpsGenieEmptyApiKey(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_opsgenie_empty_api_key"
		createEndpoint.Type = endpoints.EndpointTypeOpsGenie
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_CreateEndpointOpsGenieDuplicationError(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_opsgenie_dup"
		createEndpoint.Type = endpoints.EndpointTypeOpsGenie
		createEndpoint.ApiKey = "someApiKey"
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

func TestIntegrationEndpoints_CreateEndpointOpsGenieNoTitle(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = ""
		createEndpoint.Type = endpoints.EndpointTypeOpsGenie
		createEndpoint.ApiKey = "someApiKey"
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_UpdateEndpointOpsGenie(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_opsgenie_to_update"
		createEndpoint.Type = endpoints.EndpointTypeOpsGenie
		createEndpoint.ApiKey = "someApiKey"
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
			time.Sleep(time.Second * 1)
			createEndpoint.Title = "updated_opsgenie"
			createEndpoint.Description = "This is an UPDATED description"
			createEndpoint.ApiKey = "updatedApiKey"
			updated, err := underTest.UpdateEndpoint(int64(endpoint.Id), createEndpoint)
			assert.NoError(t, err)
			assert.NotNil(t, updated)
			assert.Equal(t, endpoint.Id, updated.Id)
		}
	}
}

func TestIntegrationEndpoints_GetEndpointOpsGenie(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_get_opsgenie"
		createEndpoint.Type = endpoints.EndpointTypeOpsGenie
		createEndpoint.ApiKey = "someApiKey"
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
			assert.Equal(t, createEndpoint.ApiKey, endpointFromGet.ApiKey)
		}
	}
}
