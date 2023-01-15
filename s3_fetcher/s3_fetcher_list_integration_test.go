package s3_fetcher_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationS3Fetcher_ListS3Fetcher(t *testing.T) {
	underTest, err := setupS3FetcherIntegrationTest()

	if assert.NoError(t, err) {
		fetchers, err := underTest.ListS3Fetchers()
		assert.NoError(t, err)
		assert.NotNil(t, fetchers)
	}
}

func TestIntegrationS3Fetcher_ListS3FetcherAtLeastOne(t *testing.T) {
	underTest, err := setupS3FetcherIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createFetcher := getCreateOrUpdateS3Fetcher(keys, false)
		fetcher, err := underTest.CreateS3Fetcher(createFetcher)
		if assert.NoError(t, err) && assert.NotNil(t, fetcher) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteS3Fetcher(fetcher.Id)
			fetchers, err := underTest.ListS3Fetchers()
			assert.NoError(t, err)
			assert.NotNil(t, fetchers)
			assert.NotEqual(t, 0, len(fetchers))
		}
	}
}
