package s3_buckets_connector_test

import (
	"encoding/json"
	"github.com/logzio/logzio_terraform_client/s3_buckets_connector"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

func TestS3BucketConnector_UpdateConnector(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	assert.NoError(t, err)
	defer teardown()

	connectorId := int64(1234567)

	mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(connectorId), 10))
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target s3_buckets_connector.S3BucketConnectorRequest
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
	})

	updateConnector := getCreateOrUpdateS3BucketConnector(keys, true)
	err = underTest.UpdateS3BucketConnector(connectorId, updateConnector)
	assert.NoError(t, err)
}

func TestS3BucketConnector_UpdateConnectorInternalServerError(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	assert.NoError(t, err)
	defer teardown()

	connectorId := int64(1234567)

	mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(connectorId), 10))
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target s3_buckets_connector.S3BucketConnectorRequest
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		w.WriteHeader(http.StatusInternalServerError)
	})

	updateConnector := getCreateOrUpdateS3BucketConnector(keys, true)
	err = underTest.UpdateS3BucketConnector(connectorId, updateConnector)
	assert.Error(t, err)
}

func TestS3BucketConnector_UpdateConnectorBucketNoPermission(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	assert.NoError(t, err)
	defer teardown()

	connectorId := int64(1234567)

	mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(connectorId), 10))
		jsonBytes, _ := ioutil.ReadAll(r.Body)
		var target s3_buckets_connector.S3BucketConnectorRequest
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		w.WriteHeader(http.StatusForbidden)
	})

	updateConnector := getCreateOrUpdateS3BucketConnector(keys, true)
	err = underTest.UpdateS3BucketConnector(connectorId, updateConnector)
	assert.Error(t, err)
}

func TestS3BucketConnector_UpdateConnectorNoBucket(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	assert.NoError(t, err)
	defer teardown()

	connectorId := int64(1234567)
	updateConnector := getCreateOrUpdateS3BucketConnector(keys, true)
	updateConnector.Bucket = ""
	err = underTest.UpdateS3BucketConnector(connectorId, updateConnector)
	assert.Error(t, err)
	assert.Equal(t, "field bucket must be set", err.Error())
}

func TestS3BucketConnector_UpdateConnectorNoRegion(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	assert.NoError(t, err)
	defer teardown()

	connectorId := int64(1234567)
	updateConnector := getCreateOrUpdateS3BucketConnector(keys, true)
	updateConnector.Region = ""
	err = underTest.UpdateS3BucketConnector(connectorId, updateConnector)
	assert.Error(t, err)
	assert.Equal(t, "field region must be set", err.Error())
}

func TestS3BucketConnector_UpdateConnectorInvalidRegion(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	assert.NoError(t, err)
	defer teardown()

	connectorId := int64(1234567)
	updateConnector := getCreateOrUpdateS3BucketConnector(keys, true)
	updateConnector.Region = "some region"
	err = underTest.UpdateS3BucketConnector(connectorId, updateConnector)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid region. region must be one of:")
}

func TestS3BucketConnector_UpdateConnectorNoLogsType(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	assert.NoError(t, err)
	defer teardown()

	connectorId := int64(1234567)
	updateConnector := getCreateOrUpdateS3BucketConnector(keys, true)
	updateConnector.LogsType = ""
	err = underTest.UpdateS3BucketConnector(connectorId, updateConnector)
	assert.Error(t, err)
	assert.Equal(t, "field logsType must be set", err.Error())
}

func TestS3BucketConnector_UpdateConnectorInvalidLogsType(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	assert.NoError(t, err)
	defer teardown()

	connectorId := int64(1234567)
	updateConnector := getCreateOrUpdateS3BucketConnector(keys, true)
	updateConnector.LogsType = "some type"
	err = underTest.UpdateS3BucketConnector(connectorId, updateConnector)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid logs type. logs type must be one of:")
}

func TestS3BucketConnector_UpdateConnectorMissingActive(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	assert.NoError(t, err)
	defer teardown()

	connectorId := int64(1234567)
	updateConnector := getCreateOrUpdateS3BucketConnector(keys, true)
	updateConnector.Active = nil
	err = underTest.UpdateS3BucketConnector(connectorId, updateConnector)
	assert.Error(t, err)
	assert.Equal(t, "field active must be set", err.Error())
}

func TestS3BucketConnector_UpdateConnectorNoAwsAuth(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	assert.NoError(t, err)
	defer teardown()

	connectorId := int64(1234567)
	updateConnector := getCreateOrUpdateS3BucketConnector(arn, true)
	updateConnector.Arn = ""
	err = underTest.UpdateS3BucketConnector(connectorId, updateConnector)
	assert.Error(t, err)
	assert.Equal(t, "either keys or arn must be set", err.Error())
}

func TestS3BucketConnector_UpdateConnectorMissingKey(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	assert.NoError(t, err)
	defer teardown()

	connectorId := int64(1234567)
	updateConnector := getCreateOrUpdateS3BucketConnector(keys, true)
	updateConnector.AccessKey = ""
	err = underTest.UpdateS3BucketConnector(connectorId, updateConnector)
	assert.Error(t, err)
	assert.Equal(t, "both aws keys must be set", err.Error())
}
