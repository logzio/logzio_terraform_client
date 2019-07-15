// +build integration

package endpoints_test

import (
	"github.com/jonboydell/logzio_client/endpoints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationEndpoints_PagerDutyCreateUpdate(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()

	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(createPagerDutyEndpoint())
		assert.NotNil(t, endpoint)

		if assert.NoError(t, err) {
			endpoint, err = underTest.UpdateEndpoint(endpoint.Id, createPagerDutyEndpoint())
			assert.NotNil(t, endpoint)
			assert.NoError(t, err)
		}
		underTest.DeleteEndpoint(endpoint.Id)
	}
}

func createPagerDutyEndpoint() endpoints.Endpoint {
	return endpoints.Endpoint{
		Title:        "pagerdutyvalidEndpoint",
		Description:  "my description",
		EndpointType: "pager-duty",
		ServiceKey:   "my_service_key",
	}
}
