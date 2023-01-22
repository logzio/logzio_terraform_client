package s3_fetcher_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationS3Fetcher_DeleteFetcher(t *testing.T) {
	underTest, err := setupS3FetcherIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		createFetcher := getCreateOrUpdateS3Fetcher(keys, false)
		fetcher, err := underTest.CreateS3Fetcher(createFetcher)
		if assert.NoError(t, err) && assert.NotNil(t, fetcher) && assert.NotZero(t, fetcher.Id) {
			time.Sleep(2 * time.Second)
			defer func() {
				err = underTest.DeleteS3Fetcher(fetcher.Id)
				assert.NoError(t, err)
			}()
		}
	}
}

func TestIntegrationS3Fetcher_DeleteFetcherNotExists(t *testing.T) {
	underTest, err := setupS3FetcherIntegrationTest()
	if assert.NoError(t, err) {
		err = underTest.DeleteS3Fetcher(1234567)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed with missing s3 fetcher")
	}
}
