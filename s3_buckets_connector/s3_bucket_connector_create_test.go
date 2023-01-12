package s3_buckets_connector_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/s3_buckets_connector"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestS3BucketsConnector_CreateS3BucketConnectorKeys(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/log-shipping/s3-buckets", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target s3_buckets_connector.S3BucketConnectorRequest
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
			fmt.Fprint(w, fixture("create_s3_bucket_connector_keys_res.json"))
		})

		createS3Connector := getCreateOrUpdateS3BucketConnector(keys, true)
		connector, err := underTest.CreateS3BucketConnector(createS3Connector)
		assert.NoError(t, err)
		assert.NotNil(t, connector)
		assert.Equal(t, int64(12345), connector.Id)
		assert.Equal(t, "my_access_key", connector.AccessKey)
		assert.Equal(t, s3_buckets_connector.LogsTypeElb, connector.LogsType)
		assert.Equal(t, s3_buckets_connector.RegionUsEast1, connector.Region)
		assert.Equal(t, "my_bucket", connector.Bucket)
		assert.Equal(t, "AWSLogs/987654321/elasticloadbalancing/us-east-1/", connector.Prefix)
		assert.True(t, connector.Active)
		assert.True(t, connector.AddS3ObjectKeyAsLogField)
	}
}

func TestS3BucketsConnector_CreateS3BucketConnectorArn(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/log-shipping/s3-buckets", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target s3_buckets_connector.S3BucketConnectorRequest
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
			fmt.Fprint(w, fixture("create_s3_bucket_connector_arn_res.json"))
		})

		createS3Connector := getCreateOrUpdateS3BucketConnector(arn, true)
		connector, err := underTest.CreateS3BucketConnector(createS3Connector)
		assert.NoError(t, err)
		assert.NotNil(t, connector)
		assert.Equal(t, int64(45678), connector.Id)
		assert.Equal(t, "my_arn", connector.Arn)
		assert.Equal(t, s3_buckets_connector.LogsTypeElb, connector.LogsType)
		assert.Equal(t, s3_buckets_connector.RegionUsEast1, connector.Region)
		assert.Equal(t, "my_bucket", connector.Bucket)
		assert.Equal(t, "AWSLogs/987654321/elasticloadbalancing/us-east-1/", connector.Prefix)
		assert.True(t, connector.Active)
		assert.True(t, connector.AddS3ObjectKeyAsLogField)
	}
}

func TestS3BucketsConnector_CreateS3BucketConnectorInternalServerError(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/log-shipping/s3-buckets", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target s3_buckets_connector.S3BucketConnectorRequest
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

		createS3Connector := getCreateOrUpdateS3BucketConnector(arn, true)
		connector, err := underTest.CreateS3BucketConnector(createS3Connector)
		assert.Error(t, err)
		assert.Nil(t, connector)
	}
}

func TestS3BucketsConnector_CreateS3BucketConnectorNoAwsAuth(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	defer teardown()

	if assert.NoError(t, err) {
		createS3Connector := getCreateOrUpdateS3BucketConnector(arn, true)
		createS3Connector.Arn = ""
		connector, err := underTest.CreateS3BucketConnector(createS3Connector)
		assert.Error(t, err)
		assert.Nil(t, connector)
		assert.Equal(t, "either keys or arn must be set", err.Error())
	}
}

func TestS3BucketsConnector_CreateS3BucketConnectorNoFullKeys(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	defer teardown()

	if assert.NoError(t, err) {
		createS3Connector := getCreateOrUpdateS3BucketConnector(keys, true)
		createS3Connector.SecretKey = ""
		connector, err := underTest.CreateS3BucketConnector(createS3Connector)
		assert.Error(t, err)
		assert.Nil(t, connector)
		assert.Equal(t, "both aws keys must be set", err.Error())
	}
}

func TestS3BucketsConnector_CreateS3BucketConnectorNoBucketName(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	defer teardown()

	if assert.NoError(t, err) {
		createS3Connector := getCreateOrUpdateS3BucketConnector(keys, true)
		createS3Connector.Bucket = ""
		connector, err := underTest.CreateS3BucketConnector(createS3Connector)
		assert.Error(t, err)
		assert.Nil(t, connector)
		assert.Equal(t, "field bucket must be set", err.Error())
	}
}

func TestS3BucketsConnector_CreateS3BucketConnectorNoRegion(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	defer teardown()

	if assert.NoError(t, err) {
		createS3Connector := getCreateOrUpdateS3BucketConnector(keys, true)
		createS3Connector.Region = ""
		connector, err := underTest.CreateS3BucketConnector(createS3Connector)
		assert.Error(t, err)
		assert.Nil(t, connector)
		assert.Equal(t, "field region must be set", err.Error())
	}
}

func TestS3BucketsConnector_CreateS3BucketConnectorInvalidRegion(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	defer teardown()

	if assert.NoError(t, err) {
		createS3Connector := getCreateOrUpdateS3BucketConnector(keys, true)
		createS3Connector.Region = "some region"
		connector, err := underTest.CreateS3BucketConnector(createS3Connector)
		assert.Error(t, err)
		assert.Nil(t, connector)
		assert.Contains(t, err.Error(), "invalid region. region must be one of:")
	}
}

func TestS3BucketsConnector_CreateS3BucketConnectorNoActiveField(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	defer teardown()

	if assert.NoError(t, err) {
		createS3Connector := getCreateOrUpdateS3BucketConnector(keys, true)
		createS3Connector.Active = nil
		connector, err := underTest.CreateS3BucketConnector(createS3Connector)
		assert.Error(t, err)
		assert.Nil(t, connector)
		assert.Equal(t, "field active must be set", err.Error())
	}
}

func TestS3BucketsConnector_CreateS3BucketConnectorNoLogsType(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	defer teardown()

	if assert.NoError(t, err) {
		createS3Connector := getCreateOrUpdateS3BucketConnector(keys, true)
		createS3Connector.LogsType = ""
		connector, err := underTest.CreateS3BucketConnector(createS3Connector)
		assert.Error(t, err)
		assert.Nil(t, connector)
		assert.Equal(t, "field logsType must be set", err.Error())
	}
}

func TestS3BucketsConnector_CreateS3BucketConnectorInvalidLogsType(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	defer teardown()

	if assert.NoError(t, err) {
		createS3Connector := getCreateOrUpdateS3BucketConnector(keys, true)
		createS3Connector.LogsType = "some type"
		connector, err := underTest.CreateS3BucketConnector(createS3Connector)
		assert.Error(t, err)
		assert.Nil(t, connector)
		assert.Contains(t, err.Error(), "invalid logs type. logs type must be one of")
	}
}
