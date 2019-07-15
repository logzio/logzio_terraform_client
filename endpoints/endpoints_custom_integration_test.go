// +build integration

package endpoints_test

import (
	"github.com/jonboydell/logzio_client/endpoints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpointsCustomCreateUpdate(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(endpoints.Endpoint{
			Title:        "testCreateCustomEndpoint",
			Method:       "POST",
			Description:  "my description",
			Url:          "https://this.is.com/some/other/webhook",
			EndpointType: "custom",
			Headers:      map[string]string{"hello": "there", "header": "two"},
			BodyTemplate: map[string]string{"hello": "there", "header": "two"},
		})
		if assert.NoError(t, err) {
			endpoint, err = underTest.UpdateEndpoint(endpoint.Id, endpoints.Endpoint{
				Title:        "testCreateUpdateCustomEndpoint",
				Method:       "POST",
				Description:  "my description update",
				Url:          "https://this.is.com/some/updated/webhook",
				EndpointType: "custom",
				Headers:      map[string]string{"hello": "there", "header": "two"},
				BodyTemplate: map[string]string{"hello": "there", "header": "two"},
			})
			assert.NotNil(t, endpoint)
			assert.NoError(t, err)
		}
		err = underTest.DeleteEndpoint(endpoint.Id)
		assert.NoError(t, err)
	}
}

func TestEndpointsCustomCreateDuplicate(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(endpoints.Endpoint{
			Title:        "testCustomDuplicateEndpoint",
			Method:       "POST",
			Description:  "my description",
			Url:          "https://this.is.com/some/other/webhook",
			EndpointType: "custom",
			Headers:      map[string]string{"hello": "there", "header": "two"},
			BodyTemplate: map[string]string{"hello": "there", "header": "two"},
		})
		if assert.NoError(t, err) {
			duplicate, err := underTest.CreateEndpoint(endpoints.Endpoint{
				Title:        "testCustomDuplicateEndpoint",
				Method:       "POST",
				Description:  "my description",
				Url:          "https://this.is.com/some/other/webhook",
				EndpointType: "custom",
				Headers:      map[string]string{"hello": "there", "header": "two"},
				BodyTemplate: map[string]string{"hello": "there", "header": "two"},
			})
			assert.Nil(t, duplicate)
			assert.Error(t, err)
		}
		err = underTest.DeleteEndpoint(endpoint.Id)
		assert.NoError(t, err)
	}
}
