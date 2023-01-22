package s3_fetcher_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationS3Fetcher_CreateS3FetcherKeys(t *testing.T) {
	underTest, err := setupS3FetcherIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createFetcher := getCreateOrUpdateS3Fetcher(keys, false)
		fetcher, err := underTest.CreateS3Fetcher(createFetcher)
		if assert.NoError(t, err) && assert.NotNil(t, fetcher) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteS3Fetcher(fetcher.Id)
			assert.NotZero(t, fetcher.Id)
		}
	}
}

func TestIntegrationS3Fetcher_CreateS3FetcherArn(t *testing.T) {
	underTest, err := setupS3FetcherIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createFetcher := getCreateOrUpdateS3Fetcher(arn, false)
		fetcher, err := underTest.CreateS3Fetcher(createFetcher)
		if assert.NoError(t, err) && assert.NotNil(t, fetcher) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteS3Fetcher(fetcher.Id)
			assert.NotZero(t, fetcher.Id)
		}
	}
}

func TestIntegrationS3Fetcher_CreateS3FetcherInvalidKeys(t *testing.T) {
	underTest, err := setupS3FetcherIntegrationTest()

	if assert.NoError(t, err) {
		createFetcher := getCreateOrUpdateS3Fetcher(keys, false)
		createFetcher.AccessKey = "some_key"
		createFetcher.SecretKey = "some_key"
		fetcher, err := underTest.CreateS3Fetcher(createFetcher)
		assert.Error(t, err)
		assert.Nil(t, fetcher)
	}
}

func TestIntegrationS3Fetcher_CreateS3FetcherInvalidArn(t *testing.T) {
	underTest, err := setupS3FetcherIntegrationTest()

	if assert.NoError(t, err) {
		createFetcher := getCreateOrUpdateS3Fetcher(arn, false)
		createFetcher.Arn = "some_arn"
		fetcher, err := underTest.CreateS3Fetcher(createFetcher)
		assert.Error(t, err)
		assert.Nil(t, fetcher)
	}
}
