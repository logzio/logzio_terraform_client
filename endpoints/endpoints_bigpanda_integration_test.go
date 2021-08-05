package endpoints_test

import (
	"github.com/logzio/logzio_terraform_client/endpoints"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestIntegrationEndpoints_CreateEndpointBigPanda(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_bigpanda"
		createEndpoint.Type = endpoints.EndpointTypeBigPanda
		createEndpoint.ApiToken = "someApiToken"
		createEndpoint.AppKey = "someAppKey"
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			time.Sleep(time.Second * 1)
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
		}
	}
}

func TestIntegrationEndpoints_CreateEndpointBigPandaNoApiToken(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_bigpanda_no_api_token"
		createEndpoint.Type = endpoints.EndpointTypeBigPanda
		createEndpoint.AppKey = "someAppKey"
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_CreateEndpointBigPandaNoAppKey(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_bigpanda_no_app_key"
		createEndpoint.Type = endpoints.EndpointTypeBigPanda
		createEndpoint.ApiToken = "someApiToken"
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_CreateEndpointBigPandaDuplicationError(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_bigpanda_dup"
		createEndpoint.Type = endpoints.EndpointTypeBigPanda
		createEndpoint.ApiToken = "someApiToken"
		createEndpoint.AppKey = "someAppKey"
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
			time.Sleep(time.Second * 1)
			duplication, err := underTest.CreateEndpoint(createEndpoint)
			assert.Error(t, err)
			assert.Nil(t, duplication)
		}
	}
}

func TestIntegrationEndpoints_CreateEndpointBigPandaNoTitle(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = ""
		createEndpoint.Type = endpoints.EndpointTypeBigPanda
		createEndpoint.ApiToken = "someApiToken"
		createEndpoint.AppKey = "someAppKey"
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_UpdateEndpointBigPanda(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_bigpanda_to_update"
		createEndpoint.Type = endpoints.EndpointTypeBigPanda
		createEndpoint.ApiToken = "someApiToken"
		createEndpoint.AppKey = "someAppKey"
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
			time.Sleep(time.Second * 1)
			createEndpoint.Title = createEndpoint.Title + "_update_bigpanda"
			createEndpoint.Description = "This is an UPDATED description"
			createEndpoint.ApiToken = "updatedApiToken"
			updated, err := underTest.UpdateEndpoint(int64(endpoint.Id), createEndpoint)
			assert.NoError(t, err)
			assert.NotNil(t, updated)
			assert.Equal(t, endpoint.Id, updated.Id)
		}
	}
}

func TestIntegrationEndpoints_GetEndpointBigPanda(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_get_bigpanda"
		createEndpoint.Type = endpoints.EndpointTypeBigPanda
		createEndpoint.ApiToken = "someApiToken"
		createEndpoint.AppKey = "someAppKey"
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			time.Sleep(time.Second * 1)
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
			endpointFromGet, err := underTest.GetEndpoint(int64(endpoint.Id))
			assert.NoError(t, err)
			assert.NotNil(t, endpointFromGet)
			assert.Equal(t, endpoint.Id, endpointFromGet.Id)
			assert.Equal(t, strings.ToLower(createEndpoint.Type), strings.ToLower(endpointFromGet.Type))
			assert.Equal(t, createEndpoint.ApiToken, endpointFromGet.ApiToken)
			assert.Equal(t, createEndpoint.AppKey, endpointFromGet.AppKey)
		}
	}
}