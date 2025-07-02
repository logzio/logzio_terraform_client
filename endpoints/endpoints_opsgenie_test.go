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

func TestEndpoints_CreateEndpointOpsGenie(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	mux.HandleFunc("/v1/endpoints/ops-genie", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, _ := io.ReadAll(r.Body)
		var target endpoints.CreateOrUpdateEndpoint
		err = json.Unmarshal(jsonBytes, &target)
		assert.NotZero(t, len(target.Title))
		assert.NotZero(t, len(target.ApiKey))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_endpoint.json"))
	})

	createEndpoint := GetCreateOrUpdateEndpoint()
	createEndpoint.Title = createEndpoint.Title + "_create_opsgenie"
	createEndpoint.ApiKey = "someApiKey"
	createEndpoint.Type = endpoints.EndpointTypeOpsGenie
	endpoint, err := underTest.CreateEndpoint(createEndpoint)
	assert.NoError(t, err)
	assert.NotNil(t, endpoint)
}

func TestEndpoints_CreateEndpointOpsGenieApiFail(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	mux.HandleFunc("/v1/endpoints/ops-genie", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, _ := io.ReadAll(r.Body)
		var target endpoints.CreateOrUpdateEndpoint
		err = json.Unmarshal(jsonBytes, &target)
		assert.NotZero(t, len(target.Title))
		assert.NotZero(t, len(target.ApiKey))
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fixture("create_endpoint_api_fail.txt"))
	})

	createEndpoint := GetCreateOrUpdateEndpoint()
	createEndpoint.Title = createEndpoint.Title + "_create_opsgenie"
	createEndpoint.ApiKey = "someApiKey"
	createEndpoint.Type = endpoints.EndpointTypeOpsGenie
	endpoint, err := underTest.CreateEndpoint(createEndpoint)
	assert.Error(t, err)
	assert.Nil(t, endpoint)
}

func TestEndpoints_UpdateEndpointOpsGenie(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	endpointId := int64(1234567)
	mux.HandleFunc("/v1/endpoints/ops-genie/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(endpointId, 10))
		jsonBytes, _ := io.ReadAll(r.Body)
		var target endpoints.CreateOrUpdateEndpoint
		err = json.Unmarshal(jsonBytes, &target)
		assert.NotZero(t, len(target.Title))
		assert.NotZero(t, len(target.ApiKey))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_endpoint.json"))
	})

	createEndpoint := GetCreateOrUpdateEndpoint()
	createEndpoint.Title = createEndpoint.Title + "_update_opsgenie"
	createEndpoint.ApiKey = "someApiKey"
	createEndpoint.Type = endpoints.EndpointTypeOpsGenie
	endpoint, err := underTest.UpdateEndpoint(endpointId, createEndpoint)
	assert.NoError(t, err)
	assert.NotNil(t, endpoint)
	assert.Equal(t, endpointId, int64(endpoint.Id))
}

func TestEndpoints_GetEndpointOpsGenie(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	endpointId := int64(1234567)
	mux.HandleFunc("/v1/endpoints/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(endpointId, 10))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_endpoint_opsgenie.json"))
	})

	endpoint, err := underTest.GetEndpoint(endpointId)
	assert.NoError(t, err)
	assert.NotNil(t, endpoint)
	assert.Equal(t, endpointId, int64(endpoint.Id))
	assert.Equal(t, endpoints.EndpointTypeOpsGenie, strings.ToLower(endpoint.Type))
	assert.Equal(t, "OpsGenie endpoint", endpoint.Title)
	assert.Equal(t, "Sends notifications to OpsGenie", endpoint.Description)
	assert.Equal(t, "someApiKey", endpoint.ApiKey)
}
