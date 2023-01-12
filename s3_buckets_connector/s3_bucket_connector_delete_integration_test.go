package s3_buckets_connector_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationS3BucketConnector_DeleteConnector(t *testing.T) {
	underTest, err := setupS3BucketConnectorIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		createConnector := getCreateOrUpdateS3BucketConnector(keys, false)
		connector, err := underTest.CreateS3BucketConnector(createConnector)
		if assert.NoError(t, err) && assert.NotNil(t, connector) && assert.NotZero(t, connector.Id) {
			time.Sleep(2 * time.Second)
			defer func() {
				err = underTest.DeleteS3BucketConnector(connector.Id)
				assert.NoError(t, err)
			}()
		}
	}
}

func TestIntegrationS3BucketConnector_DeleteConnectorNotExists(t *testing.T) {
	underTest, err := setupS3BucketConnectorIntegrationTest()
	if assert.NoError(t, err) {
		err = underTest.DeleteS3BucketConnector(1234567)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed with missing s3 bucket connector")
	}
}
