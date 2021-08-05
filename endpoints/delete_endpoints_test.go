package endpoints_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestEndpoints_DeleteEndpoint(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()
	id := int64(1234567)

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/endpoints/", func(w http.ResponseWriter, r *http.Request) {
			assert.Contains(t, r.URL.String(), strconv.FormatInt(id, 10))
			assert.Equal(t, http.MethodDelete, r.Method)
			w.WriteHeader(http.StatusNoContent)
		})

		err = underTest.DeleteEndpoint(id)
		assert.NoError(t, err)
	}
}

func TestEndpoints_DeleteEndpointIdNotFound(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()
	id := int64(1234567)

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/endpoints/", func(w http.ResponseWriter, r *http.Request) {
			assert.Contains(t, r.URL.String(), strconv.FormatInt(id, 10))
			assert.Equal(t, http.MethodDelete, r.Method)
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, fixture("delete_endpoint_not_exist.json"))
		})

		err = underTest.DeleteEndpoint(id)
		assert.Error(t, err)
	}
}

func TestEndpoints_DeleteEndpointApiFail(t *testing.T) {
	underTest, err, teardown := setupEndpointsTest()
	defer teardown()

	id := int64(1234567)
	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/endpoints/", func(w http.ResponseWriter, r *http.Request) {
			assert.Contains(t, r.URL.String(), strconv.FormatInt(id, 10))
			assert.Equal(t, http.MethodDelete, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
		})

		err = underTest.DeleteEndpoint(id)
		assert.Error(t, err)
	}
}