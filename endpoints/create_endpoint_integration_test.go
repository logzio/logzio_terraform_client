package endpoints_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestIntegrationEndpoints_CreateEndpointNoType(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_no_type"
		createEndpoint.Url = testsUrl
		createEndpoint.Method = http.MethodPost
		createEndpoint.Headers = "hello=there,header=two"
		createEndpoint.BodyTemplate = map[string]string{"hello": "there", "header": "two"}
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestIntegrationEndpoints_CreateEndpointCaseInsensitivity(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_create_case_insensitive"
		createEndpoint.Url = testsUrl
		createEndpoint.Method = http.MethodPost
		createEndpoint.Headers = "hello=there,header=two"
		createEndpoint.Type = "Custom"
		createEndpoint.BodyTemplate = map[string]string{"hello": "there", "header": "two"}
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		assert.NoError(t, err)
		assert.NotNil(t, endpoint)
		assert.NotEmpty(t, endpoint.Id)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			time.Sleep(time.Second * 1)
			defer underTest.DeleteEndpoint(int64(endpoint.Id))
		}
	}
}
