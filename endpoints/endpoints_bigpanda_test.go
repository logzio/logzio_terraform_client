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

func TestEndpoints_CreateEndpointBigPanda(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	mux.HandleFunc("/v1/endpoints/big-panda", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, _ := io.ReadAll(r.Body)
		var target endpoints.CreateOrUpdateEndpoint
		err = json.Unmarshal(jsonBytes, &target)
		assert.NotZero(t, len(target.Title))
		assert.NotZero(t, len(target.ApiToken))
		assert.NotZero(t, len(target.AppKey))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_endpoint.json"))
	})

	createEndpoint := GetCreateOrUpdateEndpoint()
	createEndpoint.Title = createEndpoint.Title + "_create_bigpanda"
	createEndpoint.ApiToken = "someApiToken"
	createEndpoint.AppKey = "someAppKey"
	createEndpoint.Type = endpoints.EndpointTypeBigPanda
	endpoint, err := underTest.CreateEndpoint(createEndpoint)
	assert.NoError(t, err)
	assert.NotNil(t, endpoint)
}

func TestEndpoints_CreateEndpointBigPandaApiFail(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	mux.HandleFunc("/v1/endpoints/big-panda", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, _ := io.ReadAll(r.Body)
		var target endpoints.CreateOrUpdateEndpoint
		err = json.Unmarshal(jsonBytes, &target)
		assert.NotZero(t, len(target.Title))
		assert.NotZero(t, len(target.ApiToken))
		assert.NotZero(t, len(target.AppKey))
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fixture("create_endpoint_api_fail.txt"))
	})

	createEndpoint := GetCreateOrUpdateEndpoint()
	createEndpoint.Title = createEndpoint.Title + "_create_bigpanda"
	createEndpoint.ApiToken = "someApiToken"
	createEndpoint.AppKey = "someAppKey"
	createEndpoint.Type = endpoints.EndpointTypeBigPanda
	endpoint, err := underTest.CreateEndpoint(createEndpoint)
	assert.Error(t, err)
	assert.Nil(t, endpoint)
}

func TestEndpoints_UpdateEndpointBigPanda(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	endpointId := int64(1234567)
	mux.HandleFunc("/v1/endpoints/big-panda/", func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.String(), strconv.FormatInt(endpointId, 10))
		assert.Equal(t, http.MethodPut, r.Method)
		jsonBytes, _ := io.ReadAll(r.Body)
		var target endpoints.CreateOrUpdateEndpoint
		err = json.Unmarshal(jsonBytes, &target)
		assert.NotZero(t, len(target.Title))
		assert.NotZero(t, len(target.ApiToken))
		assert.NotZero(t, len(target.AppKey))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_endpoint.json"))
	})

	updateEndpoint := GetCreateOrUpdateEndpoint()
	updateEndpoint.Title = updateEndpoint.Title + "_update_bigpanda"
	updateEndpoint.ApiToken = "someApiToken"
	updateEndpoint.AppKey = "someAppKey"
	updateEndpoint.Type = endpoints.EndpointTypeBigPanda
	endpoint, err := underTest.UpdateEndpoint(endpointId, updateEndpoint)
	assert.NoError(t, err)
	assert.NotNil(t, endpoint)
	assert.Equal(t, endpointId, int64(endpoint.Id))
}

func TestEndpoints_GetEndpointBigPanda(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	id := int64(1234567)
	mux.HandleFunc("/v1/endpoints/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(id, 10))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_endpoint_bigpanda.json"))
	})

	endpoint, err := underTest.GetEndpoint(id)
	assert.NoError(t, err)
	assert.NotNil(t, endpoint)
	assert.Equal(t, id, int64(endpoint.Id))
	assert.Equal(t, endpoints.EndpointTypeBigPanda, strings.ToLower(endpoint.Type))
	assert.Equal(t, "BigPanda endpoint", endpoint.Title)
	assert.Equal(t, "Sends notifications to BigPanda", endpoint.Description)
	assert.Equal(t, "someApiToken", endpoint.ApiToken)
	assert.Equal(t, "someAppKey", endpoint.AppKey)
}
