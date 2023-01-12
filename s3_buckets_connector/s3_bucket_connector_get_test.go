package s3_buckets_connector_test

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/s3_buckets_connector"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestS3BucketsConnector_GetConnectorKeys(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	assert.NoError(t, err)
	defer teardown()

	connectorId := int64(12345)

	mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(connectorId), 10))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_s3_bucket_connector_keys_res.json"))
	})

	connector, err := underTest.GetS3BucketConnector(connectorId)
	assert.NoError(t, err)
	assert.NotNil(t, connector)
	assert.Equal(t, connectorId, connector.Id)
	assert.Equal(t, "my_access_key", connector.AccessKey)
	assert.Equal(t, s3_buckets_connector.LogsTypeElb, connector.LogsType)
	assert.Equal(t, "my_bucket", connector.Bucket)
	assert.Equal(t, "AWSLogs/987654321/elasticloadbalancing/us-east-1/", connector.Prefix)
	assert.True(t, connector.Active)
	assert.Equal(t, s3_buckets_connector.RegionUsEast1, connector.Region)
	assert.Empty(t, connector.Arn)
	assert.True(t, connector.AddS3ObjectKeyAsLogField)
}

func TestS3BucketsConnector_GetConnectorArn(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	assert.NoError(t, err)
	defer teardown()

	connectorId := int64(45678)

	mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(connectorId), 10))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_s3_bucket_connector_arn_res.json"))
	})

	connector, err := underTest.GetS3BucketConnector(connectorId)
	assert.NoError(t, err)
	assert.NotNil(t, connector)
	assert.Equal(t, connectorId, connector.Id)
	assert.Equal(t, "my_arn", connector.Arn)
	assert.Equal(t, s3_buckets_connector.LogsTypeElb, connector.LogsType)
	assert.Equal(t, "my_bucket", connector.Bucket)
	assert.Equal(t, "AWSLogs/987654321/elasticloadbalancing/us-east-1/", connector.Prefix)
	assert.True(t, connector.Active)
	assert.Equal(t, s3_buckets_connector.RegionUsEast1, connector.Region)
	assert.Empty(t, connector.AccessKey)
	assert.True(t, connector.AddS3ObjectKeyAsLogField)
}

func TestS3BucketsConnector_GetConnectorInternalError(t *testing.T) {
	underTest, err, teardown := setupS3BucketConnectorTest()
	assert.NoError(t, err)
	defer teardown()

	connectorId := int64(45678)

	mux.HandleFunc("/v1/log-shipping/s3-buckets/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(connectorId), 10))
		w.WriteHeader(http.StatusInternalServerError)
	})

	connector, err := underTest.GetS3BucketConnector(connectorId)
	assert.Error(t, err)
	assert.Nil(t, connector)
}
