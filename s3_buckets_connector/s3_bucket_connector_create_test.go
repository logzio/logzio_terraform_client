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

func TestS3BucketsConnector_CreateS3BucketConnector(t *testing.T) {
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
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, fixture("create_s3_bucket_connector_res.json"))
		})

		createS3Connector := getCreateOrUpdateS3BucketConnector(keys)
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
