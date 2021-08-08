package endpoints_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestEndpoints_ListEndpoints(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	mux.HandleFunc("/v1/endpoints", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("list_endpoints.json"))
	})

	endpoints, err := underTest.ListEndpoints()
	assert.NoError(t, err)
	assert.NotNil(t, endpoints)
	assert.Equal(t, 2, len(endpoints))
}

func TestEndpoints_ListEndpointsApiFail(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	mux.HandleFunc("/v1/endpoints", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
	})

	endpoints, err := underTest.ListEndpoints()
	assert.Error(t, err)
	assert.Nil(t, endpoints)
}
