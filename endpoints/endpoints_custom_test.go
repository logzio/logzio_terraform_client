package endpoints_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/endpoints"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestEndpoints_CreateCustomEndpoint(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	mux.HandleFunc("/v1/endpoints/custom", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, _ := io.ReadAll(r.Body)
		var target endpoints.CreateOrUpdateEndpoint
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		assert.NotZero(t, len(target.Title))
		assert.NotZero(t, len(target.Url))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_endpoint.json"))
	})

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
		assert.NoError(t, err)
		assert.NotNil(t, endpoint)
	}
}

func TestEndpoints_CreateCustomEndpointApiFail(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	mux.HandleFunc("/v1/endpoints/custom", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, _ := io.ReadAll(r.Body)
		var target endpoints.CreateOrUpdateEndpoint
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		assert.NotZero(t, len(target.Title))
		assert.NotZero(t, len(target.Url))
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fixture("create_endpoint_api_fail.txt"))
	})

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
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestEndpoints_UpdateCustomEndpoint(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	endpointId := int64(1234567)
	mux.HandleFunc("/v1/endpoints/custom/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(endpointId, 10))
		jsonBytes, _ := io.ReadAll(r.Body)
		var target endpoints.CreateOrUpdateEndpoint
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		assert.NotZero(t, len(target.Title))
		assert.NotZero(t, len(target.Url))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_endpoint.json"))
	})

	if assert.NoError(t, err) {
		updateEndpoint := GetCreateOrUpdateEndpoint()
		updateEndpoint.Title = updateEndpoint.Title + "_update_custom"
		updateEndpoint.Type = endpoints.EndpointTypeCustom
		updateEndpoint.Url = testsUrl
		updateEndpoint.Method = http.MethodPost
		updateEndpoint.Headers = new(string)
		*updateEndpoint.Headers = "hello=there,header=two"
		updateEndpoint.BodyTemplate = map[string]string{"hello": "there", "header": "two"}
		endpoint, err := underTest.UpdateEndpoint(endpointId, updateEndpoint)
		assert.NoError(t, err)
		assert.NotNil(t, endpoint)
		assert.Equal(t, endpointId, int64(endpoint.Id))
	}
}

func TestEndpoints_GetCustomEndpoint(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	endpointId := int64(1234567)
	mux.HandleFunc("/v1/endpoints/", func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.String(), strconv.FormatInt(endpointId, 10))
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_endpoint_custom.json"))
	})

	if assert.NoError(t, err) {
		endpoint, err := underTest.GetEndpoint(endpointId)
		assert.NoError(t, err)
		assert.NotNil(t, endpoint)
		assert.Equal(t, endpointId, int64(endpoint.Id))
		assert.Equal(t, endpoints.EndpointTypeCustom, strings.ToLower(endpoint.Type))
		assert.Equal(t, "New custom endpoint", endpoint.Title)
		assert.Equal(t, "Sends notifications to my custom endpoint", endpoint.Description)
		assert.Equal(t, testsUrl, endpoint.Url)
		assert.Equal(t, http.MethodPost, endpoint.Method)
		assert.Equal(t, "authKey=6e30-60a9-3591", endpoint.Headers)
		assert.NotEmpty(t, endpoint.BodyTemplate)
	}
}
