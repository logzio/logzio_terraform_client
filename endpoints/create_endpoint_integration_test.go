package endpoints_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
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
