package s3_buckets_connector_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestIntegrationS3BucketConnector_CreateS3BucketConnectorKeys(t *testing.T) {
	underTest, err := setupS3BucketConnectorIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createConnector := getCreateOrUpdateS3BucketConnector(keys, false)
		connector, err := underTest.CreateS3BucketConnector(createConnector)
		if assert.NoError(t, err) && assert.NotNil(t, connector) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteS3BucketConnector(connector.Id)
			assert.NotZero(t, connector.Id)
		}
	}
}

func TestIntegrationS3BucketConnector_CreateS3BucketConnectorArn(t *testing.T) {
	underTest, err := setupS3BucketConnectorIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createConnector := getCreateOrUpdateS3BucketConnector(arn, false)
		connector, err := underTest.CreateS3BucketConnector(createConnector)
		if assert.NoError(t, err) && assert.NotNil(t, connector) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteS3BucketConnector(connector.Id)
			assert.NotZero(t, connector.Id)
		}
	}
}

func TestIntegrationS3BucketConnector_CreateS3BucketConnectorInvalidKeys(t *testing.T) {
	underTest, err := setupS3BucketConnectorIntegrationTest()

	if assert.NoError(t, err) {
		createConnector := getCreateOrUpdateS3BucketConnector(keys, false)
		createConnector.AccessKey = "some_key"
		createConnector.SecretKey = "some_key"
		connector, err := underTest.CreateS3BucketConnector(createConnector)
		assert.Error(t, err)
		assert.Nil(t, connector)
	}
}

func TestIntegrationS3BucketConnector_CreateS3BucketConnectorInvalidArn(t *testing.T) {
	underTest, err := setupS3BucketConnectorIntegrationTest()

	if assert.NoError(t, err) {
		createConnector := getCreateOrUpdateS3BucketConnector(arn, false)
		createConnector.Arn = "some_arn"
		connector, err := underTest.CreateS3BucketConnector(createConnector)
		assert.Error(t, err)
		assert.Nil(t, connector)
	}
}

func TestIntegrationS3BucketConnector_CreateS3BucketConnectorArnWithoutPermissions(t *testing.T) {
	underTest, err := setupS3BucketConnectorIntegrationTest()

	if assert.NoError(t, err) {
		createConnector := getCreateOrUpdateS3BucketConnector(arn, false)
		createConnector.Arn = os.Getenv("AWS_ARN")
		connector, err := underTest.CreateS3BucketConnector(createConnector)
		assert.Error(t, err)
		assert.Nil(t, connector)
	}
}

func TestIntegrationS3BucketConnector_CreateS3BucketConnectorInvalidBucket(t *testing.T) {
	underTest, err := setupS3BucketConnectorIntegrationTest()

	if assert.NoError(t, err) {
		createConnector := getCreateOrUpdateS3BucketConnector(keys, false)
		createConnector.Bucket = "some_bucket"
		connector, err := underTest.CreateS3BucketConnector(createConnector)
		assert.Error(t, err)
		assert.Nil(t, connector)
	}
}