package s3_buckets_connector_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestS3BucketConnector_ListConnectors(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("list_s3_bucket_connector.json"))
	})

	connectors, err := underTest.ListS3BucketConnectors()
	assert.NoError(t, err)
	assert.NotNil(t, connectors)
	assert.Equal(t, 2, len(connectors))
}

func TestS3BucketConnector_ListConnectorsInternalServer(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
	})

	connectors, err := underTest.ListS3BucketConnectors()
	assert.Error(t, err)
	assert.Nil(t, connectors)
}
