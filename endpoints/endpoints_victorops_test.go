package endpoints_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/endpoints"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestEndpoints_CreateEndpointVictorOps(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	mux.HandleFunc("/v1/endpoints/victorops", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target endpoints.CreateOrUpdateEndpoint
		err = json.Unmarshal(jsonBytes, &target)
		assert.NotZero(t, len(target.Title))
		assert.NotZero(t, len(target.RoutingKey))
		assert.NotZero(t, len(target.MessageType))
		assert.NotZero(t, len(target.ServiceApiKey))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_endpoint.json"))
	})

	createEndpoint := GetCreateOrUpdateEndpoint()
	createEndpoint.Title = createEndpoint.Title + "_create_victorops"
	createEndpoint.RoutingKey = "someRoutingKey"
	createEndpoint.MessageType = "someMessageType"
	createEndpoint.ServiceApiKey = "someServiceApiKey"
	createEndpoint.Type = endpoints.EndpointTypeVictorOps
	endpoint, err := underTest.CreateEndpoint(createEndpoint)
	assert.NoError(t, err)
	assert.NotNil(t, endpoint)
}

func TestEndpoints_CreateEndpointVictorOpsApiFail(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	mux.HandleFunc("/v1/endpoints/victorops", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target endpoints.CreateOrUpdateEndpoint
		err = json.Unmarshal(jsonBytes, &target)
		assert.NotZero(t, len(target.Title))
		assert.NotZero(t, len(target.RoutingKey))
		assert.NotZero(t, len(target.MessageType))
		assert.NotZero(t, len(target.ServiceApiKey))
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fixture("create_endpoint_api_fail.txt"))
	})

	createEndpoint := GetCreateOrUpdateEndpoint()
	createEndpoint.Title = createEndpoint.Title + "_create_victorops"
	createEndpoint.RoutingKey = "someRoutingKey"
	createEndpoint.MessageType = "someMessageType"
	createEndpoint.ServiceApiKey = "someServiceApiKey"
	createEndpoint.Type = endpoints.EndpointTypeVictorOps
	endpoint, err := underTest.CreateEndpoint(createEndpoint)
	assert.Error(t, err)
	assert.Nil(t, endpoint)
}

func TestEndpoints_UpdateEndpointVictorOps(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	endpointId := int64(1234567)
	mux.HandleFunc("/v1/endpoints/victorops/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(endpointId, 10))
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target endpoints.CreateOrUpdateEndpoint
		err = json.Unmarshal(jsonBytes, &target)
		assert.NotZero(t, len(target.Title))
		assert.NotZero(t, len(target.RoutingKey))
		assert.NotZero(t, len(target.MessageType))
		assert.NotZero(t, len(target.ServiceApiKey))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_endpoint.json"))
	})

	createEndpoint := GetCreateOrUpdateEndpoint()
	createEndpoint.Title = createEndpoint.Title + "_update_victorops"
	createEndpoint.RoutingKey = "someRoutingKey"
	createEndpoint.MessageType = "someMessageType"
	createEndpoint.ServiceApiKey = "someServiceApiKey"
	createEndpoint.Type = endpoints.EndpointTypeVictorOps
	endpoint, err := underTest.UpdateEndpoint(endpointId, createEndpoint)
	assert.NoError(t, err)
	assert.NotNil(t, endpoint)
	assert.Equal(t, endpointId, int64(endpoint.Id))
}

func TestEndpoints_GetEndpointVictorOps(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	endpointId := int64(1234567)
	mux.HandleFunc("/v1/endpoints/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(endpointId, 10))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_endpoint_victorops.json"))
	})

	endpoint, err := underTest.GetEndpoint(endpointId)
	assert.NoError(t, err)
	assert.NotNil(t, endpoint)
	assert.Equal(t, endpointId, int64(endpoint.Id))
	assert.Equal(t, endpoints.EndpointTypeVictorOps, strings.ToLower(endpoint.Type))
	assert.Equal(t, "VictorOps endpoint", endpoint.Title)
	assert.Equal(t, "Sends notifications to VictorOps", endpoint.Description)
	assert.Equal(t, "someRoutingKey", endpoint.RoutingKey)
	assert.Equal(t, "someMessageType", endpoint.MessageType)
	assert.Equal(t, "someServiceApiKey", endpoint.ServiceApiKey)
}