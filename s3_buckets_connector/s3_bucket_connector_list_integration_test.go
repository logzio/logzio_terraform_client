package s3_buckets_connector_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationS3BucketConnector_ListS3BucketConnector(t *testing.T) {
	underTest, err := setupS3BucketConnectorIntegrationTest()

	if assert.NoError(t, err) {
		connectors, err := underTest.ListS3BucketConnectors()
		assert.NoError(t, err)
		assert.NotNil(t, connectors)
	}
}

func TestIntegrationS3BucketConnector_ListS3BucketConnectorAtLeastOne(t *testing.T) {
	underTest, err := setupS3BucketConnectorIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createConnector := getCreateOrUpdateS3BucketConnector(keys, false)
		connector, err := underTest.CreateS3BucketConnector(createConnector)
		if assert.NoError(t, err) && assert.NotNil(t, connector) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteS3BucketConnector(connector.Id)
			connectors, err := underTest.ListS3BucketConnectors()
			assert.NoError(t, err)
			assert.NotNil(t, connectors)
			assert.NotEqual(t, 0, len(connectors))
		}
	}
}
