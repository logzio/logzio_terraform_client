package s3_buckets_connector_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestS3BucketsConnector_DeleteConnector(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	defer teardown()

	if assert.NoError(t, err) {
		connectorId := int64(1234567)

		mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(connectorId, 10))
			w.WriteHeader(http.StatusOK)
		})

		err = underTest.DeleteS3BucketConnector(connectorId)
		assert.NoError(t, err)
	}
}

func TestS3BucketsConnector_DeleteConnectorInternalError(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	defer teardown()

	if assert.NoError(t, err) {
		connectorId := int64(1234567)

		mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(connectorId, 10))
			w.WriteHeader(http.StatusInternalServerError)
		})

		err = underTest.DeleteS3BucketConnector(connectorId)
		assert.Error(t, err)
	}
}
