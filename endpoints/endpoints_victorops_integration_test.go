package endpoints_test

import (
	"github.com/logzio/logzio_terraform_client/endpoints"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestIntegrationEndpoints_CreateEndpointVicrorOps(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_victorops"
		createEndpoint.Type = endpoints.EndpointTypeVictorOps
		createEndpoint.RoutingKey = "someRoutingKey"
		createEndpoint.MessageType = "someMessageType"
		createEndpoint.ServiceApiKey = "someServiceApiKey"
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
			time.Sleep(time.Second * 1)
		}
	}
}

func TestIntegrationEndpoints_CreateEndpointVicrorOpsEmptyRoutingKey(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_victorops_empty_routing_key"
		createEndpoint.Type = endpoints.EndpointTypeVictorOps
		createEndpoint.MessageType = "someMessageType"
		createEndpoint.ServiceApiKey = "someServiceApiKey"
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_CreateEndpointVicrorOpsEmptyMessageType(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_victorops_empty_message_type"
		createEndpoint.Type = endpoints.EndpointTypeVictorOps
		createEndpoint.RoutingKey = "someRoutingKey"
		createEndpoint.ServiceApiKey = "someServiceApiKey"
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_CreateEndpointVicrorOpsEmptyServiceApiKey(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_victorops_empty_service_api_key"
		createEndpoint.Type = endpoints.EndpointTypeVictorOps
		createEndpoint.RoutingKey = "someRoutingKey"
		createEndpoint.MessageType = "someMessageType"
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_CreateEndpointVicrorOpsDuplicationError(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_victorops_dup"
		createEndpoint.Type = endpoints.EndpointTypeVictorOps
		createEndpoint.RoutingKey = "someRoutingKey"
		createEndpoint.MessageType = "someMessageType"
		createEndpoint.ServiceApiKey = "someServiceApiKey"
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

func TestIntegrationEndpoints_CreateEndpointVicrorOpsNoTitle(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = ""
		createEndpoint.Type = endpoints.EndpointTypeVictorOps
		createEndpoint.RoutingKey = "someRoutingKey"
		createEndpoint.MessageType = "someMessageType"
		createEndpoint.ServiceApiKey = "someServiceApiKey"
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_UpdateEndpointVictorOps(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_victorops_to_update"
		createEndpoint.Type = endpoints.EndpointTypeVictorOps
		createEndpoint.RoutingKey = "someRoutingKey"
		createEndpoint.MessageType = "someMessageType"
		createEndpoint.ServiceApiKey = "someServiceApiKey"
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
			time.Sleep(time.Second * 1)
			createEndpoint.Title = "updated_victorops"
			createEndpoint.Description = "This is an UPDATED description"
			createEndpoint.RoutingKey = "updatedRoutingKey"
			createEndpoint.MessageType = "updatedMessageType"
			createEndpoint.ServiceApiKey = "updatedServiceApiKey"
			updated, err := underTest.UpdateEndpoint(int64(endpoint.Id), createEndpoint)
			assert.NoError(t, err)
			assert.NotNil(t, updated)
			assert.Equal(t, endpoint.Id, updated.Id)
		}
	}
}

func TestIntegrationEndpoints_GetEndpointVictorOps(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_get_victorops"
		createEndpoint.Type = endpoints.EndpointTypeVictorOps
		createEndpoint.RoutingKey = "someRoutingKey"
		createEndpoint.MessageType = "someMessageType"
		createEndpoint.ServiceApiKey = "someServiceApiKey"
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
			assert.Equal(t, createEndpoint.RoutingKey, endpointFromGet.RoutingKey)
			assert.Equal(t, createEndpoint.MessageType, endpointFromGet.MessageType)
			assert.Equal(t, createEndpoint.ServiceApiKey, endpointFromGet.ServiceApiKey)
		}
	}
}
