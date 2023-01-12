package s3_buckets_connector_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationS3BucketConnector_GetConnectorKeys(t *testing.T) {
	underTest, err := setupS3BucketConnectorIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createConnector := getCreateOrUpdateS3BucketConnector(keys, false)
		connector, err := underTest.CreateS3BucketConnector(createConnector)
		if assert.NoError(t, err) && assert.NotNil(t, connector) && assert.NotZero(t, connector.Id) {
			defer underTest.DeleteS3BucketConnector(connector.Id)
			time.Sleep(4 * time.Second)
			getConnector, err := underTest.GetS3BucketConnector(connector.Id)
			assert.NoError(t, err)
			assert.NotNil(t, getConnector)
			assert.Equal(t, connector, getConnector)
		}
	}
}

func TestIntegrationS3BucketConnector_GetConnectorArn(t *testing.T) {
	underTest, err := setupS3BucketConnectorIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createConnector := getCreateOrUpdateS3BucketConnector(arn, false)
		connector, err := underTest.CreateS3BucketConnector(createConnector)
		if assert.NoError(t, err) && assert.NotNil(t, connector) && assert.NotZero(t, connector.Id) {
			defer underTest.DeleteS3BucketConnector(connector.Id)
			time.Sleep(4 * time.Second)
			getConnector, err := underTest.GetS3BucketConnector(connector.Id)
			assert.NoError(t, err)
			assert.NotNil(t, getConnector)
			assert.Equal(t, connector, getConnector)
		}
	}
}

func TestIntegrationS3BucketConnector_GetConnectorIdNotExists(t *testing.T) {
	underTest, err := setupS3BucketConnectorIntegrationTest()

	if assert.NoError(t, err) {
		connector, err := underTest.GetS3BucketConnector(int64(123456))
		assert.Error(t, err)
		assert.Nil(t, connector)
		assert.Contains(t, err.Error(), "failed with missing s3 bucket connector")
	}
}
