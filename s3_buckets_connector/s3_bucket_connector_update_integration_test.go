package s3_buckets_connector_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationS3BucketConnector_UpdateConnector(t *testing.T) {
	underTest, err := setupS3BucketConnectorIntegrationTest()

	if assert.NoError(t, err) {
		request := getCreateOrUpdateS3BucketConnector(keys, false)

		createConnector, err := underTest.CreateS3BucketConnector(request)
		if assert.NoError(t, err) && assert.NotNil(t, createConnector) && assert.NotZero(t, createConnector.Id) {
			defer underTest.DeleteS3BucketConnector(createConnector.Id)
			time.Sleep(time.Second * 2)
			active := false
			request.Active = &active
			err = underTest.UpdateS3BucketConnector(createConnector.Id, request)
			assert.NoError(t, err)
			// verify that the update was made
			time.Sleep(time.Second * 2)
			getConnector, err := underTest.GetS3BucketConnector(createConnector.Id)
			assert.NoError(t, err)
			assert.False(t, getConnector.Active)
		}
	}
}

func TestIntegrationS3BucketConnector_UpdateConnectorIdNotFound(t *testing.T) {
	underTest, err := setupS3BucketConnectorIntegrationTest()

	if assert.NoError(t, err) {
		request := getCreateOrUpdateS3BucketConnector(keys, false)
		err = underTest.UpdateS3BucketConnector(int64(123456), request)
		assert.Error(t, err)
	}
}
