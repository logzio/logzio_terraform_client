package endpoints_test

import (
	"github.com/logzio/logzio_terraform_client/endpoints"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestIntegrationEndpoints_CreateEndpointCustom(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_custom"
		createEndpoint.Type = endpoints.EndpointTypeCustom
		createEndpoint.Url = testsUrl
		createEndpoint.Method = http.MethodPost
		createEndpoint.Headers = new(string)
		*createEndpoint.Headers = "hello=there,header=two"
		createEndpoint.BodyTemplate = map[string]string{"hello": "there", "header": "two"}
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			time.Sleep(time.Second * 1)
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
		}
	}
}

func TestIntegrationEndpoints_CreateEndpointCustomEmptyHeaders(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_custom_empty_headers"
		createEndpoint.Type = endpoints.EndpointTypeCustom
		createEndpoint.Url = testsUrl
		createEndpoint.Method = http.MethodPost
		createEndpoint.BodyTemplate = map[string]string{"hello": "there", "header": "two"}
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
		}
	}
}

func TestIntegrationEndpoints_CreateEndpointCustomEmptyBodyTemplate(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_custom_empty_body_template"
		createEndpoint.Type = endpoints.EndpointTypeCustom
		createEndpoint.Url = testsUrl
		createEndpoint.Method = http.MethodPost
		createEndpoint.Headers = new(string)
		*createEndpoint.Headers = "hello=there,header=two"
		createEndpoint.BodyTemplate = nil
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
		}
	}
}

func TestIntegrationEndpoints_CreateEndpointCustomNoUrl(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_custom_no_url"
		createEndpoint.Type = endpoints.EndpointTypeCustom
		createEndpoint.Method = http.MethodPost
		createEndpoint.Headers = new(string)
		*createEndpoint.Headers = "hello=there,header=two"
		createEndpoint.BodyTemplate = map[string]string{"hello": "there", "header": "two"}
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_CreateEndpointCustomNoMethod(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_custom_no_method"
		createEndpoint.Type = endpoints.EndpointTypeCustom
		createEndpoint.Url = testsUrl
		createEndpoint.Headers = new(string)
		*createEndpoint.Headers = "hello=there,header=two"
		createEndpoint.BodyTemplate = map[string]string{"hello": "there", "header": "two"}
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_CreateEndpointCustomDuplicationError(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_custom_dup"
		createEndpoint.Type = endpoints.EndpointTypeCustom
		createEndpoint.Url = testsUrl
		createEndpoint.Method = http.MethodPost
		createEndpoint.Headers = new(string)
		*createEndpoint.Headers = "hello=there,header=two"
		createEndpoint.BodyTemplate = map[string]string{"hello": "there", "header": "two"}
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
			time.Sleep(time.Second * 1)
			duplicated, err := underTest.CreateEndpoint(createEndpoint)
			assert.Error(t, err)
			assert.Nil(t, duplicated)
		}
	}
}

func TestIntegrationEndpoints_CreateEndpointNoTitle(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = ""
		createEndpoint.Type = endpoints.EndpointTypeCustom
		createEndpoint.Url = testsUrl
		createEndpoint.Method = http.MethodPost
		createEndpoint.Headers = new(string)
		*createEndpoint.Headers = "hello=there,header=two"
		createEndpoint.BodyTemplate = map[string]string{"hello": "there", "header": "two"}
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_UpdateEndpointCustom(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_custom_to_update"
		createEndpoint.Type = endpoints.EndpointTypeCustom
		createEndpoint.Url = testsUrl
		createEndpoint.Method = http.MethodPost
		createEndpoint.Headers = new(string)
		*createEndpoint.Headers = "hello=there,header=two"
		createEndpoint.BodyTemplate = map[string]string{"hello": "there", "header": "two"}
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
			time.Sleep(time.Second * 1)
			createEndpoint.Title = "updated_custom"
			createEndpoint.Description = "This is an UPDATED description"
			createEndpoint.Method = http.MethodPut
			createEndpoint.Url = testsUrlUpdate
			*createEndpoint.Headers = ""
			updated, err := underTest.UpdateEndpoint(int64(endpoint.Id), createEndpoint)
			assert.NoError(t, err)
			assert.NotNil(t, updated)
			assert.Equal(t, endpoint.Id, updated.Id)
			time.Sleep(2 * time.Second)
			updatedGet, err := underTest.GetEndpoint(int64(updated.Id))
			assert.NoError(t, err)
			assert.NotNil(t, updatedGet)
			assert.Equal(t, createEndpoint.Title, updatedGet.Title)
			assert.Equal(t, createEndpoint.Description, updatedGet.Description)
			assert.Equal(t, createEndpoint.Method, updatedGet.Method)
			assert.Equal(t, createEndpoint.Url, updatedGet.Url)
			assert.Equal(t, *createEndpoint.Headers, updatedGet.Headers)
		}
	}
}

func TestIntegrationEndpoints_GetEndpointCustom(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_get_custom"
		createEndpoint.Type = endpoints.EndpointTypeCustom
		createEndpoint.Url = testsUrl
		createEndpoint.Method = http.MethodPost
		createEndpoint.Headers = new(string)
		*createEndpoint.Headers = "hello=there,header=two"
		createEndpoint.BodyTemplate = map[string]string{"hello": "there", "header": "two"}
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			time.Sleep(time.Second * 1)
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
			endpointFromGet, err := underTest.GetEndpoint(int64(endpoint.Id))
			assert.NoError(t, err)
			assert.NotNil(t, endpointFromGet)
			assert.Equal(t, endpoint.Id, endpointFromGet.Id)
			assert.Equal(t, createEndpoint.Title, endpointFromGet.Title)
			assert.Equal(t, strings.ToLower(createEndpoint.Type), strings.ToLower(endpointFromGet.Type))
			assert.Equal(t, createEndpoint.Url, endpointFromGet.Url)
			assert.Equal(t, createEndpoint.Method, endpointFromGet.Method)
			assert.Equal(t, *createEndpoint.Headers, endpointFromGet.Headers)
			assert.NotEmpty(t, endpointFromGet.BodyTemplate)
		}
	}
}
