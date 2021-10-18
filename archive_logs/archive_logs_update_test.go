package archive_logs_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/archive_logs"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

func TestArchiveLogs_UpdateArchiveS3(t *testing.T) {
	underTest, teardown, err := setupArchiveLogsTest()
	defer teardown()
	id := int32(1234)

	if assert.NoError(t, err) {
		mux.HandleFunc(archiveApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target archive_logs.CreateOrUpdateArchiving
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(id), 10))
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("update_archive_s3.json"))
		})

		createArchive, err := test_utils.GetCreateOrUpdateArchiveLogs(archive_logs.StorageTypeS3)
		if assert.NoError(t, err) {
			archive, err := underTest.UpdateArchiveLogs(id, createArchive)
			assert.NoError(t, err)
			assert.NotNil(t, archive)
			assert.Equal(t, int32(1234), archive.Id)
			assert.Equal(t, archive_logs.StorageTypeS3, archive.Settings.StorageType)
			assert.True(t, archive.Settings.Enabled)
			assert.True(t, archive.Settings.Compressed)
			assert.Equal(t, archive_logs.CredentialsTypeKeys, archive.Settings.AmazonS3StorageSettings.CredentialsType)
			assert.Equal(t, "another-path", archive.Settings.AmazonS3StorageSettings.Path)
			assert.Equal(t, "another-access-key", archive.Settings.AmazonS3StorageSettings.S3SecretCredentials.AccessKey)
			assert.Equal(t, "another-secret-key", archive.Settings.AmazonS3StorageSettings.S3SecretCredentials.SecretKey)
			assert.Empty(t, archive.Settings.AmazonS3StorageSettings.S3IamCredentials)
			assert.Empty(t, archive.Settings.AzureBlobStorageSettings)
		}
	}
}

func TestArchiveLogs_UpdateArchiveBlob(t *testing.T) {
	underTest, teardown, err := setupArchiveLogsTest()
	defer teardown()
	id := int32(1234)

	if assert.NoError(t, err) {
		mux.HandleFunc(archiveApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target archive_logs.CreateOrUpdateArchiving
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(id), 10))
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("update_archive_blob.json"))
		})

		createArchive, err := test_utils.GetCreateOrUpdateArchiveLogs(archive_logs.StorageTypeBlob)
		if assert.NoError(t, err) {
			createArchive.AzureBlobStorageSettings.TenantId = "another-tenant-id"
			createArchive.AzureBlobStorageSettings.ClientId = "another-client-id"
			createArchive.AzureBlobStorageSettings.ClientSecret = "another-client-secret"
			createArchive.AzureBlobStorageSettings.AccountName = "another-account-name"
			createArchive.AzureBlobStorageSettings.ContainerName = "another-container-name"
			archive, err := underTest.UpdateArchiveLogs(id, createArchive)
			assert.NoError(t, err)
			assert.NotNil(t, archive)
			assert.Equal(t, id, archive.Id)
			assert.Equal(t, archive_logs.StorageTypeBlob, archive.Settings.StorageType)
			assert.True(t, archive.Settings.Enabled)
			assert.True(t, archive.Settings.Compressed)
			assert.Equal(t, "another-tenant-id", archive.Settings.AzureBlobStorageSettings.TenantId)
			assert.Equal(t, "another-client-id", archive.Settings.AzureBlobStorageSettings.ClientId)
			assert.Equal(t, "another-client-secret", archive.Settings.AzureBlobStorageSettings.ClientSecret)
			assert.Equal(t, "another-account-name", archive.Settings.AzureBlobStorageSettings.AccountName)
			assert.Equal(t, "another-container-name", archive.Settings.AzureBlobStorageSettings.ContainerName)
			assert.Empty(t, archive.Settings.AmazonS3StorageSettings)
		}
	}
}

func TestArchiveLogs_UpdateArchiveIdNotFound(t *testing.T) {
	underTest, teardown, err := setupArchiveLogsTest()
	defer teardown()
	id := int32(1234)

	if assert.NoError(t, err) {
		mux.HandleFunc(archiveApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target archive_logs.CreateOrUpdateArchiving
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(id), 10))
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, fixture("update_archive_id_not_found.txt"))
		})

		createArchive, err := test_utils.GetCreateOrUpdateArchiveLogs(archive_logs.StorageTypeBlob)
		if assert.NoError(t, err) {
			archive, err := underTest.UpdateArchiveLogs(id, createArchive)
			assert.Error(t, err)
			assert.Nil(t, archive)
		}
	}
}

func TestArchiveLogs_UpdateArchiveApiFail(t *testing.T) {
	underTest, teardown, err := setupArchiveLogsTest()
	defer teardown()

	mux.HandleFunc(archiveApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fixture("archive_api_fail.txt"))
	})

	if assert.NoError(t, err) {
		createArchive, err := test_utils.GetCreateOrUpdateArchiveLogs(archive_logs.StorageTypeS3)
		if assert.NoError(t, err) {
			archive, err := underTest.UpdateArchiveLogs(int32(1234), createArchive)
			assert.Error(t, err)
			assert.Nil(t, archive)
		}
	}
}
