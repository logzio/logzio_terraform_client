package endpoints_test

import (
	"github.com/logzio/logzio_terraform_client/endpoints"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationEndpoints_DeleteEndpoint(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()

	if assert.NoError(t, err) {
		createEndpoint := GetCreateOrUpdateEndpoint()
		createEndpoint.Title = createEndpoint.Title + "_delete_test"
		createEndpoint.Type = endpoints.EndpointTypeSlack
		createEndpoint.Url = testsUrl
		endpoint, err := underTest.CreateEndpoint(createEndpoint)
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) && assert.NotEmpty(t, endpoint.Id) {
			time.Sleep(time.Second * 1)
			defer func() {
				err = underTest.DeleteEndpoint(int64(endpoint.Id))
				assert.NoError(t, err)
			}()

		}
	}
}

func TestIntegrationEndpoints_DeleteEndpointIdNotFound(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()

	if assert.NoError(t, err) {
		err = underTest.DeleteEndpoint(int64(1234567))
		assert.Error(t, err)
	}
}
