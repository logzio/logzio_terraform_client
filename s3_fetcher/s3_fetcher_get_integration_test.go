package s3_fetcher_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationS3Fetcher_GetFetcherKeys(t *testing.T) {
	underTest, err := setupS3FetcherIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createFetcher := getCreateOrUpdateS3Fetcher(keys, false)
		fetcher, err := underTest.CreateS3Fetcher(createFetcher)
		if assert.NoError(t, err) && assert.NotNil(t, fetcher) && assert.NotZero(t, fetcher.Id) {
			defer underTest.DeleteS3Fetcher(fetcher.Id)
			time.Sleep(4 * time.Second)
			getFetcher, err := underTest.GetS3Fetcher(fetcher.Id)
			assert.NoError(t, err)
			assert.NotNil(t, getFetcher)
			assert.Equal(t, fetcher, getFetcher)
		}
	}
}

func TestIntegrationS3Fetcher_GetFetcherArn(t *testing.T) {
	underTest, err := setupS3FetcherIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createFetcher := getCreateOrUpdateS3Fetcher(arn, false)
		fetcher, err := underTest.CreateS3Fetcher(createFetcher)
		if assert.NoError(t, err) && assert.NotNil(t, fetcher) && assert.NotZero(t, fetcher.Id) {
			defer underTest.DeleteS3Fetcher(fetcher.Id)
			time.Sleep(4 * time.Second)
			getFetcher, err := underTest.GetS3Fetcher(fetcher.Id)
			assert.NoError(t, err)
			assert.NotNil(t, getFetcher)
			assert.Equal(t, fetcher, getFetcher)
		}
	}
}

func TestIntegrationS3Fetcher_GetFetcherIdNotExists(t *testing.T) {
	underTest, err := setupS3FetcherIntegrationTest()

	if assert.NoError(t, err) {
		fetcher, err := underTest.GetS3Fetcher(int64(123456))
		assert.Error(t, err)
		assert.Nil(t, fetcher)
		assert.Contains(t, err.Error(), "failed with missing s3 fetcher")
	}
}
