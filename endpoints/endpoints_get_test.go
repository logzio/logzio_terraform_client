package endpoints_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestEndpoints_GetEndpoint(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	endpointId := int64(1234567)

	mux.HandleFunc("/v1/endpoints/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(endpointId, 10))
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_endpoint.json"))
	})

	if assert.NoError(t, err) {
		endpoints, err := underTest.GetEndpoint(endpointId)
		assert.NoError(t, err)
		assert.NotNil(t, endpoints)
	}
}

func TestEndpoints_GetEndpointByName(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	mux.HandleFunc("/v1/endpoints", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("list_endpoints.json"))
	})

	if assert.NoError(t, err) {
		endpoint, err := underTest.GetEndpointByName("my endpoint")
		assert.NoError(t, err)
		assert.NotNil(t, endpoint)
		assert.Equal(t, "my endpoint", endpoint.Title)
	}
}

func TestEndpoints_GetEndpointNotExist(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	endpointId := int64(1234567)

	mux.HandleFunc("/v1/endpoints/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(endpointId, 10))
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_endpoint_not_exist.json"))
	})

	if assert.NoError(t, err) {
		endpoints, err := underTest.GetEndpoint(endpointId)
		assert.Error(t, err)
		assert.Nil(t, endpoints)
	}
}
