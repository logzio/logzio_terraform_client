package s3_fetcher_test

import (
	"encoding/json"
	"github.com/logzio/logzio_terraform_client/s3_fetcher"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

func TestS3Fetcher_UpdateFetcher(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	assert.NoError(t, err)
	defer teardown()

	fetcherId := int64(1234567)

	mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(fetcherId), 10))
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target s3_fetcher.S3FetcherRequest
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
	})

	updateFetcher := getCreateOrUpdateS3Fetcher(keys, true)
	err = underTest.UpdateS3Fetcher(fetcherId, updateFetcher)
	assert.NoError(t, err)
}

func TestS3Fetcher_UpdateFetcherInternalServerError(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	assert.NoError(t, err)
	defer teardown()

	fetcherId := int64(1234567)

	mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(fetcherId), 10))
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target s3_fetcher.S3FetcherRequest
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		w.WriteHeader(http.StatusInternalServerError)
	})

	updateFetcher := getCreateOrUpdateS3Fetcher(keys, true)
	err = underTest.UpdateS3Fetcher(fetcherId, updateFetcher)
	assert.Error(t, err)
}

func TestS3Fetcher_UpdateFetcherNoPermission(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	assert.NoError(t, err)
	defer teardown()

	fetcherId := int64(1234567)

	mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(fetcherId), 10))
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target s3_fetcher.S3FetcherRequest
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		w.WriteHeader(http.StatusForbidden)
	})

	updateFetcher := getCreateOrUpdateS3Fetcher(keys, true)
	err = underTest.UpdateS3Fetcher(fetcherId, updateFetcher)
	assert.Error(t, err)
}

func TestS3Fetcher_UpdateFetcherNoBucket(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	assert.NoError(t, err)
	defer teardown()

	fetcherId := int64(1234567)
	updateFetcher := getCreateOrUpdateS3Fetcher(keys, true)
	updateFetcher.Bucket = ""
	err = underTest.UpdateS3Fetcher(fetcherId, updateFetcher)
	assert.Error(t, err)
	assert.Equal(t, "field bucket must be set", err.Error())
}

func TestS3Fetcher_UpdateFetcherNoRegion(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	assert.NoError(t, err)
	defer teardown()

	fetcherId := int64(1234567)
	updateFetcher := getCreateOrUpdateS3Fetcher(keys, true)
	updateFetcher.Region = ""
	err = underTest.UpdateS3Fetcher(fetcherId, updateFetcher)
	assert.Error(t, err)
	assert.Equal(t, "field region must be set", err.Error())
}

func TestS3Fetcher_UpdateFetcherInvalidRegion(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	assert.NoError(t, err)
	defer teardown()

	fetcherId := int64(1234567)
	updateFetcher := getCreateOrUpdateS3Fetcher(keys, true)
	updateFetcher.Region = "some region"
	err = underTest.UpdateS3Fetcher(fetcherId, updateFetcher)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid region. region must be one of:")
}

func TestS3Fetcher_UpdateFetcherNoLogsType(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	assert.NoError(t, err)
	defer teardown()

	fetcherId := int64(1234567)
	updateFetcher := getCreateOrUpdateS3Fetcher(keys, true)
	updateFetcher.LogsType = ""
	err = underTest.UpdateS3Fetcher(fetcherId, updateFetcher)
	assert.Error(t, err)
	assert.Equal(t, "field logsType must be set", err.Error())
}

func TestS3Fetcher_UpdateFetcherInvalidLogsType(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	assert.NoError(t, err)
	defer teardown()

	fetcherId := int64(1234567)
	updateFetcher := getCreateOrUpdateS3Fetcher(keys, true)
	updateFetcher.LogsType = "some type"
	err = underTest.UpdateS3Fetcher(fetcherId, updateFetcher)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid logs type. logs type must be one of:")
}

func TestS3Fetcher_UpdateFetcherMissingActive(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	assert.NoError(t, err)
	defer teardown()

	fetcherId := int64(1234567)
	updateFetcher := getCreateOrUpdateS3Fetcher(keys, true)
	updateFetcher.Active = nil
	err = underTest.UpdateS3Fetcher(fetcherId, updateFetcher)
	assert.Error(t, err)
	assert.Equal(t, "field active must be set", err.Error())
}

func TestS3Fetcher_UpdateFetcherNoAwsAuth(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	assert.NoError(t, err)
	defer teardown()

	fetcherId := int64(1234567)
	updateFetcher := getCreateOrUpdateS3Fetcher(arn, true)
	updateFetcher.Arn = ""
	err = underTest.UpdateS3Fetcher(fetcherId, updateFetcher)
	assert.Error(t, err)
	assert.Equal(t, "either keys or arn must be set", err.Error())
}

func TestS3Fetcher_UpdateFetcherMissingKey(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	assert.NoError(t, err)
	defer teardown()

	fetcherId := int64(1234567)
	updateFetcher := getCreateOrUpdateS3Fetcher(keys, true)
	updateFetcher.AccessKey = ""
	err = underTest.UpdateS3Fetcher(fetcherId, updateFetcher)
	assert.Error(t, err)
	assert.Equal(t, "both aws keys must be set", err.Error())
}
