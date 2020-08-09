package endpoints_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/endpoints"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

func TestEndpoints_CreateSlackEndpoint(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	endpointId := int64(1234567)

	mux.HandleFunc("/v1/endpoints/slack", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.Equal(t, http.MethodPost, r.Method)
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target map[string]interface{}
		err = json.Unmarshal(jsonBytes, &target)
		assert.Contains(t, target, "title")
		assert.Contains(t, target, "description")
		assert.Contains(t, target, "url")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_endpoint.json"))
	})

	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(endpoints.Endpoint{
			Title:        "some_endpoint",
			Description:  "my description",
			Url:          "https://this.is.com/some/webhook",
			EndpointType: endpoints.EndpointTypeSlack,
		})
		assert.NoError(t, err)
		assert.NotNil(t, endpoint)
		assert.Equal(t, endpointId, endpoint.Id)
	}
}

func TestEndpoints_UpdateSlackEndpoint(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	endpointId := int64(1234567)

	mux.HandleFunc("/v1/endpoints/slack/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(endpointId, 10))
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target map[string]interface{}
		err = json.Unmarshal(jsonBytes, &target)
		assert.Contains(t, target, "title")
		assert.Contains(t, target, "description")
		assert.Contains(t, target, "url")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_endpoint.json"))
	})

	if assert.NoError(t, err) {
		endpoint, err := underTest.UpdateEndpoint(endpointId, endpoints.Endpoint{
			Title:        "some_endpoint",
			Description:  "my description",
			Url:          "https://this.is.com/some/webhook",
			EndpointType: endpoints.EndpointTypeSlack,
		})
		assert.NoError(t, err)
		assert.NotNil(t, endpoint)
		assert.Equal(t, endpointId, endpoint.Id)
	}
}

func TestEndpoints_CreateDuplicateEndpoint(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	mux.HandleFunc("/v1/endpoints/slack", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_duplicate_endpoint.json"))
	})

	if assert.NoError(t, err) {
		endpoint, err := underTest.CreateEndpoint(endpoints.Endpoint{
			Title:        "some_endpoint",
			Description:  "my description",
			Url:          "https://this.is.com/some/webhook",
			EndpointType: endpoints.EndpointTypeSlack,
		})
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}

func TestEndpoints_UpdateSlackEndpointToDuplicate(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	endpointId := int64(1234567)

	mux.HandleFunc("/v1/endpoints/slack/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(endpointId, 10))
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_endpoint_to_duplicate.json"))
	})

	if assert.NoError(t, err) {
		endpoint, err := underTest.UpdateEndpoint(endpointId, endpoints.Endpoint{
			Title:        "some_endpoint",
			Description:  "my description",
			Url:          "https://this.is.com/some/webhook",
			EndpointType: endpoints.EndpointTypeSlack,
		})
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}
