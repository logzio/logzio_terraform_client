package s3_fetcher_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/s3_fetcher"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestS3Fetcher_CreateS3FetcherKeys(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/log-shipping/s3-buckets", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target s3_fetcher.S3FetcherRequest
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Bucket)
			assert.NotNil(t, target.Active)
			assert.NotEmpty(t, target.Region)
			assert.NotEmpty(t, target.LogsType)
			assert.NotEmpty(t, target.AccessKey)
			assert.NotEmpty(t, target.SecretKey)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, fixture("create_s3_fetcher_keys_res.json"))
		})

		createS3Fetcher := getCreateOrUpdateS3Fetcher(keys, true)
		fetcher, err := underTest.CreateS3Fetcher(createS3Fetcher)
		assert.NoError(t, err)
		assert.NotNil(t, fetcher)
		assert.Equal(t, int64(12345), fetcher.Id)
		assert.Equal(t, "my_access_key", fetcher.AccessKey)
		assert.Equal(t, LogsTypeS3Access, fetcher.LogsType)
		assert.Equal(t, s3_fetcher.RegionUsEast1, fetcher.Region)
		assert.Equal(t, "my_bucket", fetcher.Bucket)
		assert.Equal(t, "AWSLogs/987654321/elasticloadbalancing/us-east-1/", fetcher.Prefix)
		assert.True(t, fetcher.Active)
		assert.True(t, fetcher.AddS3ObjectKeyAsLogField)
	}
}

func TestS3Fetcher_CreateS3FetcherArn(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/log-shipping/s3-buckets", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target s3_fetcher.S3FetcherRequest
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Bucket)
			assert.NotNil(t, target.Active)
			assert.NotEmpty(t, target.Region)
			assert.NotEmpty(t, target.LogsType)
			assert.NotEmpty(t, target.Arn)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, fixture("create_s3_fetcher_arn_res.json"))
		})

		createS3Fetcher := getCreateOrUpdateS3Fetcher(arn, true)
		fetcher, err := underTest.CreateS3Fetcher(createS3Fetcher)
		assert.NoError(t, err)
		assert.NotNil(t, fetcher)
		assert.Equal(t, int64(45678), fetcher.Id)
		assert.Equal(t, "my_arn", fetcher.Arn)
		assert.Equal(t, LogsTypeS3Access, fetcher.LogsType)
		assert.Equal(t, s3_fetcher.RegionUsEast1, fetcher.Region)
		assert.Equal(t, "my_bucket", fetcher.Bucket)
		assert.Equal(t, "AWSLogs/987654321/elasticloadbalancing/us-east-1/", fetcher.Prefix)
		assert.True(t, fetcher.Active)
		assert.True(t, fetcher.AddS3ObjectKeyAsLogField)
	}
}

func TestS3Fetcher_CreateS3FetcherInternalServerError(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/log-shipping/s3-buckets", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target s3_fetcher.S3FetcherRequest
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Bucket)
			assert.NotNil(t, target.Active)
			assert.NotEmpty(t, target.Region)
			assert.NotEmpty(t, target.LogsType)
			assert.NotEmpty(t, target.Arn)
			w.WriteHeader(http.StatusInternalServerError)
		})

		createS3Fetcher := getCreateOrUpdateS3Fetcher(arn, true)
		fetcher, err := underTest.CreateS3Fetcher(createS3Fetcher)
		assert.Error(t, err)
		assert.Nil(t, fetcher)
	}
}

func TestS3Fetcher_CreateS3FetcherNoAwsAuth(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	defer teardown()

	if assert.NoError(t, err) {
		createS3Fetcher := getCreateOrUpdateS3Fetcher(arn, true)
		createS3Fetcher.Arn = ""
		fetcher, err := underTest.CreateS3Fetcher(createS3Fetcher)
		assert.Error(t, err)
		assert.Nil(t, fetcher)
		assert.Equal(t, "either keys or arn must be set", err.Error())
	}
}

func TestS3Fetcher_CreateS3FetcherNoFullKeys(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	defer teardown()

	if assert.NoError(t, err) {
		createS3Fetcher := getCreateOrUpdateS3Fetcher(keys, true)
		createS3Fetcher.SecretKey = ""
		fetcher, err := underTest.CreateS3Fetcher(createS3Fetcher)
		assert.Error(t, err)
		assert.Nil(t, fetcher)
		assert.Equal(t, "both aws keys must be set", err.Error())
	}
}

func TestS3Fetcher_CreateS3FetcherNoBucketName(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	defer teardown()

	if assert.NoError(t, err) {
		createS3Fetcher := getCreateOrUpdateS3Fetcher(keys, true)
		createS3Fetcher.Bucket = ""
		fetcher, err := underTest.CreateS3Fetcher(createS3Fetcher)
		assert.Error(t, err)
		assert.Nil(t, fetcher)
		assert.Equal(t, "field bucket must be set", err.Error())
	}
}

func TestS3Fetcher_CreateS3FetcherNoRegion(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	defer teardown()

	if assert.NoError(t, err) {
		createS3Fetcher := getCreateOrUpdateS3Fetcher(keys, true)
		createS3Fetcher.Region = ""
		fetcher, err := underTest.CreateS3Fetcher(createS3Fetcher)
		assert.Error(t, err)
		assert.Nil(t, fetcher)
		assert.Equal(t, "field region must be set", err.Error())
	}
}

func TestS3Fetcher_CreateS3FetcherInvalidRegion(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	defer teardown()

	if assert.NoError(t, err) {
		createS3Fetcher := getCreateOrUpdateS3Fetcher(keys, true)
		createS3Fetcher.Region = "some region"
		fetcher, err := underTest.CreateS3Fetcher(createS3Fetcher)
		assert.Error(t, err)
		assert.Nil(t, fetcher)
		assert.Contains(t, err.Error(), "invalid region. region must be one of:")
	}
}

func TestS3Fetcher_CreateS3FetcherNoActiveField(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	defer teardown()

	if assert.NoError(t, err) {
		createS3Fetcher := getCreateOrUpdateS3Fetcher(keys, true)
		createS3Fetcher.Active = nil
		fetcher, err := underTest.CreateS3Fetcher(createS3Fetcher)
		assert.Error(t, err)
		assert.Nil(t, fetcher)
		assert.Equal(t, "field active must be set", err.Error())
	}
}

func TestS3Fetcher_CreateS3FetcherNoLogsType(t *testing.T) {
	underTest, err, teardown := setupS3FetcherTest()
	defer teardown()

	if assert.NoError(t, err) {
		createS3Fetcher := getCreateOrUpdateS3Fetcher(keys, true)
		createS3Fetcher.LogsType = ""
		fetcher, err := underTest.CreateS3Fetcher(createS3Fetcher)
		assert.Error(t, err)
		assert.Nil(t, fetcher)
		assert.Equal(t, "field logsType must be set", err.Error())
	}
}
