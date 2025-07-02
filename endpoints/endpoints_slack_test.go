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

func TestEndpoints_CreateEndpointSlack(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	mux.HandleFunc("/v1/endpoints/slack", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, _ := io.ReadAll(r.Body)
		var target endpoints.CreateOrUpdateEndpoint
		err = json.Unmarshal(jsonBytes, &target)
		assert.NotZero(t, len(target.Title))
		assert.NotZero(t, len(target.Url))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_endpoint.json"))
	})

	createEndpoint := GetCreateOrUpdateEndpoint()
	createEndpoint.Title = createEndpoint.Title + "_create_slack"
	createEndpoint.Url = testsUrl
	createEndpoint.Type = endpoints.EndpointTypeSlack
	endpoint, err := underTest.CreateEndpoint(createEndpoint)
	assert.NoError(t, err)
	assert.NotNil(t, endpoint)
}

func TestEndpoints_CreateEndpointSlackApiFail(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	mux.HandleFunc("/v1/endpoints/slack", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, _ := io.ReadAll(r.Body)
		var target endpoints.CreateOrUpdateEndpoint
		err = json.Unmarshal(jsonBytes, &target)
		assert.NotZero(t, len(target.Title))
		assert.NotZero(t, len(target.Url))
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fixture("create_endpoint_api_fail.txt"))
	})

	createEndpoint := GetCreateOrUpdateEndpoint()
	createEndpoint.Title = createEndpoint.Title + "_create_slack"
	createEndpoint.Url = testsUrl
	createEndpoint.Type = endpoints.EndpointTypeSlack
	endpoint, err := underTest.CreateEndpoint(createEndpoint)
	assert.Error(t, err)
	assert.Nil(t, endpoint)
}

func TestEndpoints_UpdateEndpointSlack(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	endpointId := int64(1234567)
	mux.HandleFunc("/v1/endpoints/slack/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(endpointId, 10))
		jsonBytes, _ := io.ReadAll(r.Body)
		var target endpoints.CreateOrUpdateEndpoint
		err = json.Unmarshal(jsonBytes, &target)
		assert.NotZero(t, len(target.Title))
		assert.NotZero(t, len(target.Url))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_endpoint.json"))
	})

	createEndpoint := GetCreateOrUpdateEndpoint()
	createEndpoint.Title = createEndpoint.Title + "_update_slack"
	createEndpoint.Url = testsUrl
	createEndpoint.Type = endpoints.EndpointTypeSlack
	endpoint, err := underTest.UpdateEndpoint(endpointId, createEndpoint)
	assert.NoError(t, err)
	assert.NotNil(t, endpoint)
	assert.Equal(t, endpointId, int64(endpoint.Id))
}

func TestEndpoints_GetEndpointSlack(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	endpointId := int64(1234567)
	mux.HandleFunc("/v1/endpoints/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(endpointId, 10))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_endpoint_slack.json"))
	})

	endpoint, err := underTest.GetEndpoint(endpointId)
	assert.NoError(t, err)
	assert.NotNil(t, endpoint)
	assert.Equal(t, endpointId, int64(endpoint.Id))
	assert.Equal(t, endpoints.EndpointTypeSlack, strings.ToLower(endpoint.Type))
	assert.Equal(t, "New Slack endpoint", endpoint.Title)
	assert.Equal(t, "Sends notifications to logzio-alerts channel", endpoint.Description)
	assert.Equal(t, testsUrl, endpoint.Url)
}
