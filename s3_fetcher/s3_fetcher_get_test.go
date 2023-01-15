package s3_fetcher_test

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/s3_fetcher"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestS3Fetcher_GetFetcherKeys(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	assert.NoError(t, err)
	defer teardown()

	fetcherId := int64(12345)

	mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(fetcherId), 10))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_s3_fetcher_keys_res.json"))
	})

	fetcher, err := underTest.GetS3Fetcher(fetcherId)
	assert.NoError(t, err)
	assert.NotNil(t, fetcher)
	assert.Equal(t, fetcherId, fetcher.Id)
	assert.Equal(t, "my_access_key", fetcher.AccessKey)
	assert.Equal(t, s3_fetcher.LogsTypeElb, fetcher.LogsType)
	assert.Equal(t, "my_bucket", fetcher.Bucket)
	assert.Equal(t, "AWSLogs/987654321/elasticloadbalancing/us-east-1/", fetcher.Prefix)
	assert.True(t, fetcher.Active)
	assert.Equal(t, s3_fetcher.RegionUsEast1, fetcher.Region)
	assert.Empty(t, fetcher.Arn)
	assert.True(t, fetcher.AddS3ObjectKeyAsLogField)
}

func TestS3Fetcher_GetFetcherArn(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	assert.NoError(t, err)
	defer teardown()

	fetcherId := int64(45678)

	mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(fetcherId), 10))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_s3_fetcher_arn_res.json"))
	})

	fetcher, err := underTest.GetS3Fetcher(fetcherId)
	assert.NoError(t, err)
	assert.NotNil(t, fetcher)
	assert.Equal(t, fetcherId, fetcher.Id)
	assert.Equal(t, "my_arn", fetcher.Arn)
	assert.Equal(t, s3_fetcher.LogsTypeElb, fetcher.LogsType)
	assert.Equal(t, "my_bucket", fetcher.Bucket)
	assert.Equal(t, "AWSLogs/987654321/elasticloadbalancing/us-east-1/", fetcher.Prefix)
	assert.True(t, fetcher.Active)
	assert.Equal(t, s3_fetcher.RegionUsEast1, fetcher.Region)
	assert.Empty(t, fetcher.AccessKey)
	assert.True(t, fetcher.AddS3ObjectKeyAsLogField)
}

func TestS3Fetcher_GetFetcherInternalError(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	assert.NoError(t, err)
	defer teardown()

	fetcherId := int64(45678)

	mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(fetcherId), 10))
		w.WriteHeader(http.StatusInternalServerError)
	})

	fetcher, err := underTest.GetS3Fetcher(fetcherId)
	assert.Error(t, err)
	assert.Nil(t, fetcher)
}
