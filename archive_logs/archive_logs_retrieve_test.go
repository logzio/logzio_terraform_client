package archive_logs_test

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/archive_logs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestArchiveLogs_RetrieveArchive(t *testing.T) {
	underTest, teardown, err := setupArchiveLogsTest()
	defer teardown()
	assert.NoError(t, err)
	archiveId := int32(1234)

	mux.HandleFunc(archiveApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(archiveId), 10))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_archive.json"))
	})

	archive, err := underTest.RetrieveArchiveLogsSetting(archiveId)
	assert.NoError(t, err)
	assert.NotNil(t, archive)
	assert.Equal(t, archiveId, archive.Id)
	assert.Equal(t, archive_logs.StorageTypeBlob, archive.Settings.StorageType)
	assert.True(t, archive.Settings.Enabled)
	assert.True(t, archive.Settings.Compressed)
	assert.Equal(t, "some-tenant-id", archive.Settings.AzureBlobStorageSettings.TenantId)
	assert.Equal(t, "some-client-id", archive.Settings.AzureBlobStorageSettings.ClientId)
	assert.Equal(t, "some-client-secret", archive.Settings.AzureBlobStorageSettings.ClientSecret)
	assert.Equal(t, "some-account-name", archive.Settings.AzureBlobStorageSettings.AccountName)
	assert.Equal(t, "some-container-name", archive.Settings.AzureBlobStorageSettings.ContainerName)
	assert.Empty(t, archive.Settings.AmazonS3StorageSettings)
}

func TestArchiveLogs_RetrieveArchiveApiFail(t *testing.T) {
	underTest, teardown, err := setupArchiveLogsTest()
	defer teardown()
	assert.NoError(t, err)
	archiveId := int32(1234)

	mux.HandleFunc(archiveApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(archiveId), 10))
		w.WriteHeader(http.StatusInternalServerError)
	})

	archive, err := underTest.RetrieveArchiveLogsSetting(archiveId)
	assert.Error(t, err)
	assert.Nil(t, archive)
}

func TestArchiveLogs_RetrieveArchiveIdNotFound(t *testing.T) {
	underTest, teardown, err := setupArchiveLogsTest()
	defer teardown()
	assert.NoError(t, err)
	archiveId := int32(1234)

	mux.HandleFunc(archiveApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(archiveId), 10))
		w.WriteHeader(http.StatusNotFound)
	})

	archive, err := underTest.RetrieveArchiveLogsSetting(archiveId)
	assert.Error(t, err)
	assert.Nil(t, archive)
	assert.Contains(t, err.Error(), "failed with missing archive")
}
