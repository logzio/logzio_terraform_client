package s3_fetcher_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestS3Fetcher_DeleteFetcher(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	defer teardown()

	if assert.NoError(t, err) {
		fetcherId := int64(1234567)

		mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(fetcherId, 10))
			w.WriteHeader(http.StatusOK)
		})

		err = underTest.DeleteS3Fetcher(fetcherId)
		assert.NoError(t, err)
	}
}

func TestS3Fetcher_DeleteFetcherInternalError(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	defer teardown()

	if assert.NoError(t, err) {
		fetcherId := int64(1234567)

		mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(fetcherId, 10))
			w.WriteHeader(http.StatusInternalServerError)
		})

		err = underTest.DeleteS3Fetcher(fetcherId)
		assert.Error(t, err)
	}
}
