package endpoints_test

import (
	"github.com/logzio/logzio_terraform_client/endpoints"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationEndpoints_CustomCreateUpdate(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(endpoints.Endpoint{
			Title:        "testCreateCustomEndpoint",
			Method:       "POST",
			Description:  "my description",
			Url:          "https://jsonplaceholder.typicode.com/todos/1",
			EndpointType: "custom",
			Headers:      map[string]string{"hello": "there", "header": "two"},
			BodyTemplate: map[string]string{"hello": "there", "header": "two"},
		})
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(endpoint.Id)

			endpoint, err = underTest.UpdateEndpoint(endpoint.Id, endpoints.Endpoint{
				Title:        "testCreateUpdateCustomEndpoint",
				Method:       "POST",
				Description:  "my description update",
				Url:          "https://jsonplaceholder.typicode.com/todos/1",
				EndpointType: "custom",
				Headers:      map[string]string{"hello": "there", "header": "two"},
				BodyTemplate: map[string]string{"hello": "there", "header": "two"},
			})
			assert.NotNil(t, endpoint)
			assert.NoError(t, err)
		}
	}
}

func TestIntegrationEndpoints_CustomCreateDuplicate(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(endpoints.Endpoint{
			Title:        "testCustomDuplicateEndpoint",
			Method:       "POST",
			Description:  "my description",
			Url:          "https://jsonplaceholder.typicode.com/todos/1",
			EndpointType: "custom",
			Headers:      map[string]string{"hello": "there", "header": "two"},
			BodyTemplate: map[string]string{"hello": "there", "header": "two"},
		})
		if assert.NoError(t, err) && assert.NotNil(t, endpoint) {
			defer underTest.DeleteEndpoint(endpoint.Id)

			duplicate, err := underTest.CreateEndpoint(endpoints.Endpoint{
				Title:        "testCustomDuplicateEndpoint",
				Method:       "POST",
				Description:  "my description",
				Url:          "https://jsonplaceholder.typicode.com/todos/1",
				EndpointType: "custom",
				Headers:      map[string]string{"hello": "there", "header": "two"},
				BodyTemplate: map[string]string{"hello": "there", "header": "two"},
			})
			assert.Nil(t, duplicate)
			assert.Error(t, err)
		}
	}
}

func TestIntegrationEndpoints_CustomCreateNoHeader(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(endpoints.Endpoint{
			Title:        "testCreateCustomEndpointNoHeaders",
			Method:       "POST",
			Description:  "my description",
			Url:          "https://jsonplaceholder.typicode.com/todos/1",
			EndpointType: "custom",
			BodyTemplate: map[string]string{"hello": "there", "header": "two"},
		})

		if endpoint != nil {
			defer underTest.DeleteEndpoint(endpoint.Id)
		}

		assert.NotNil(t, endpoint)
		assert.NoError(t, err)
		assert.NotEmpty(t, endpoint.Id)
	}
}

func TestIntegrationEndpoints_CustomGetNoHeaders(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(endpoints.Endpoint{
			Title:        "testCreateCustomEndpointNoHeaders",
			Method:       "POST",
			Description:  "my description",
			Url:          "https://jsonplaceholder.typicode.com/todos/1",
			EndpointType: "custom",
			BodyTemplate: map[string]string{"hello": "there", "header": "two"},
		})

		if endpoint != nil {
			defer underTest.DeleteEndpoint(endpoint.Id)
		}

		assert.NotNil(t, endpoint)
		assert.NoError(t, err)
		assert.NotEmpty(t, endpoint.Id)

		time.Sleep(4 * time.Second)

		getEndpoint, err := underTest.GetEndpoint(endpoint.Id)

		assert.NotNil(t, endpoint)
		assert.NoError(t, err)
		assert.Equal(t, endpoint.Id, getEndpoint.Id)
	}
}
