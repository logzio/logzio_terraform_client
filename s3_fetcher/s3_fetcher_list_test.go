package s3_fetcher_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestS3Fetcher_ListFetchers(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("list_s3_fetcher.json"))
	})

	fetchers, err := underTest.ListS3Fetchers()
	assert.NoError(t, err)
	assert.NotNil(t, fetchers)
	assert.Equal(t, 2, len(fetchers))
}

func TestS3Fetcher_ListFetchersInternalServer(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
	})

	fetchers, err := underTest.ListS3Fetchers()
	assert.Error(t, err)
	assert.Nil(t, fetchers)
}
