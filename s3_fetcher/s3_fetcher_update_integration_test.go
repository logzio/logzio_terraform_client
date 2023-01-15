package s3_fetcher_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationS3Fetcher_UpdateFetcher(t *testing.T) {
	underTest, err := setupS3FetcherIntegrationTest()

	if assert.NoError(t, err) {
		request := getCreateOrUpdateS3Fetcher(keys, false)

		createFetcher, err := underTest.CreateS3Fetcher(request)
		if assert.NoError(t, err) && assert.NotNil(t, createFetcher) && assert.NotZero(t, createFetcher.Id) {
			defer underTest.DeleteS3Fetcher(createFetcher.Id)
			time.Sleep(time.Second * 2)
			active := false
			request.Active = &active
			err = underTest.UpdateS3Fetcher(createFetcher.Id, request)
			assert.NoError(t, err)
			// verify that the update was made
			time.Sleep(time.Second * 2)
			getFetcher, err := underTest.GetS3Fetcher(createFetcher.Id)
			assert.NoError(t, err)
			assert.False(t, getFetcher.Active)
		}
	}
}

func TestIntegrationS3Fetcher_UpdateFetcherIdNotFound(t *testing.T) {
	underTest, err := setupS3FetcherIntegrationTest()

	if assert.NoError(t, err) {
		request := getCreateOrUpdateS3Fetcher(keys, false)
		err = underTest.UpdateS3Fetcher(int64(123456), request)
		assert.Error(t, err)
	}
}
