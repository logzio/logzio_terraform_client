package endpoints_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestEndpoints_GetEndpointNotExist(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	endpointId := int64(123456)
	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/endpoints/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("get_endpoint_not_exist.json"))
		})
	}

	endpoint, err := underTest.GetEndpoint(endpointId)
	assert.Error(t, err)
	assert.Nil(t, endpoint)
}
