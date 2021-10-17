package archive_logs_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/archive_logs"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestArchiveLogs_SetupArchiveS3Keys(t *testing.T) {
	underTest, err, teardown := setupArchiveLogsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v2/archive/settings", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target archive_logs.CreateOrUpdateArchiving
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.Equal(t, archive_logs.StorageTypeS3, target.StorageType)
			assert.NotEmpty(t, target.AmazonS3StorageSettings)
			assert.Equal(t, archive_logs.CredentialsTypeKeys, target.AmazonS3StorageSettings.CredentialsType)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("setup_archive_s3_keys.json"))
		})

		createArchive, err := test_utils.GetCreateOrUpdateArchiveLogs(archive_logs.StorageTypeS3)
		createArchive.AmazonS3StorageSettings.Path = "some-path"
		createArchive.AmazonS3StorageSettings.S3SecretCredentials.AccessKey = "some-access-key"
		createArchive.AmazonS3StorageSettings.S3SecretCredentials.SecretKey = "some-secret-key"
		if assert.NoError(t, err) {
			archive, err := underTest.SetupArchive(createArchive)
			assert.NoError(t, err)
			assert.NotNil(t, archive)
			assert.Equal(t, int32(1234), archive.Id)
			assert.Equal(t, archive_logs.StorageTypeS3, archive.Settings.StorageType)
			assert.True(t, archive.Settings.Enabled)
			assert.True(t, archive.Settings.Compressed)
			assert.Equal(t, archive_logs.CredentialsTypeKeys, archive.Settings.AmazonS3StorageSettings.CredentialsType)
			assert.Equal(t, "some-path", archive.Settings.AmazonS3StorageSettings.Path)
			assert.Equal(t, "some-acces-key", archive.Settings.AmazonS3StorageSettings.S3SecretCredentials.AccessKey)
			assert.Equal(t, "some-secret-key", archive.Settings.AmazonS3StorageSettings.S3SecretCredentials.SecretKey)
			assert.Empty(t, archive.Settings.AmazonS3StorageSettings.S3IamCredentials)
			assert.Empty(t, archive.Settings.AzureBlobStorageSettings)
		}
	}
}

func TestArchiveLogs_SetupArchiveS3Iam(t *testing.T) {
	underTest, err, teardown := setupArchiveLogsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v2/archive/settings", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target archive_logs.CreateOrUpdateArchiving
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.Equal(t, archive_logs.StorageTypeS3, target.StorageType)
			assert.NotEmpty(t, target.AmazonS3StorageSettings)
			assert.Equal(t, archive_logs.CredentialsTypeIam, target.AmazonS3StorageSettings.CredentialsType)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("setup_archive_s3_iam.json"))
		})

		createArchive, err := test_utils.GetCreateOrUpdateArchiveLogs(archive_logs.StorageTypeS3)
		createArchive.AmazonS3StorageSettings.S3SecretCredentials = nil
		createArchive.AmazonS3StorageSettings.CredentialsType = archive_logs.CredentialsTypeIam
		iamCredentials, err := test_utils.GetS3IamCredentials()
		if assert.NoError(t, err) {
			createArchive.AmazonS3StorageSettings.S3IamCredentials = iamCredentials
			createArchive.AmazonS3StorageSettings.S3IamCredentials.Arn = "some-arn"
			createArchive.AmazonS3StorageSettings.Path = "some-path"
			if assert.NoError(t, err) {
				archive, err := underTest.SetupArchive(createArchive)
				assert.NoError(t, err)
				assert.NotNil(t, archive)
				assert.Equal(t, int32(1234), archive.Id)
				assert.Equal(t, archive_logs.StorageTypeS3, archive.Settings.StorageType)
				assert.True(t, archive.Settings.Enabled)
				assert.True(t, archive.Settings.Compressed)
				assert.Equal(t, archive_logs.CredentialsTypeIam, archive.Settings.AmazonS3StorageSettings.CredentialsType)
				assert.Equal(t, "some-path", archive.Settings.AmazonS3StorageSettings.Path)
				assert.Equal(t, "some-arn", archive.Settings.AmazonS3StorageSettings.S3IamCredentials.Arn)
				assert.Equal(t, "some-external-id", archive.Settings.AmazonS3StorageSettings.S3IamCredentials.ExternalId)
				assert.Empty(t, archive.Settings.AmazonS3StorageSettings.S3SecretCredentials)
				assert.Empty(t, archive.Settings.AzureBlobStorageSettings)
			}
		}
	}
}

func TestArchiveLogs_SetupArchiveBlob(t *testing.T) {
	underTest, err, teardown := setupArchiveLogsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v2/archive/settings", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target archive_logs.CreateOrUpdateArchiving
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.Equal(t, archive_logs.StorageTypeBlob, target.StorageType)
			assert.NotEmpty(t, target.AzureBlobStorageSettings)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("setup_archive_blob.json"))
		})

		createArchive, err := test_utils.GetCreateOrUpdateArchiveLogs(archive_logs.StorageTypeBlob)
		if assert.NoError(t, err) {
			createArchive.AzureBlobStorageSettings.TenantId = "some-tenant-id"
			createArchive.AzureBlobStorageSettings.ClientId = "some-client-id"
			createArchive.AzureBlobStorageSettings.ClientSecret = "some-client-secret"
			createArchive.AzureBlobStorageSettings.AccountName = "some-account-name"
			createArchive.AzureBlobStorageSettings.ContainerName = "some-container-name"
			archive, err := underTest.SetupArchive(createArchive)
			assert.NoError(t, err)
			assert.NotNil(t, archive)
			assert.Equal(t, int32(1234), archive.Id)
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
	}
}

func TestArchiveLogs_SetupArchiveInvalidStorageType(t *testing.T) {
	underTest, err, teardown := setupArchiveLogsTest()
	defer teardown()

	createArchive, err := test_utils.GetCreateOrUpdateArchiveLogs(archive_logs.StorageTypeS3)
	if assert.NoError(t, err) {
		createArchive.StorageType = "invalid storage type"
		archive, err := underTest.SetupArchive(createArchive)
		assert.Error(t, err)
		assert.Nil(t, archive)
	}
}

func TestArchiveLogs_SetupArchiveS3InvalidCredentialsType(t *testing.T) {
	underTest, err, teardown := setupArchiveLogsTest()
	defer teardown()

	if assert.NoError(t, err) {
		createArchive, err := test_utils.GetCreateOrUpdateArchiveLogs(archive_logs.StorageTypeS3)
		if assert.NoError(t, err) {
			createArchive.AmazonS3StorageSettings.CredentialsType = "invalid credentials type"
			archive, err := underTest.SetupArchive(createArchive)
			assert.Error(t, err)
			assert.Nil(t, archive)
		}
	}
}

func TestArchiveLogs_SetupArchiveApiFail(t *testing.T) {
	underTest, err, teardown := setupArchiveLogsTest()
	defer teardown()

	mux.HandleFunc("/v2/archive/settings", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fixture("archive_api_fail.txt"))
	})

	if assert.NoError(t, err) {
		createArchive, err := test_utils.GetCreateOrUpdateArchiveLogs(archive_logs.StorageTypeS3)
		if assert.NoError(t, err) {
			archive, err := underTest.SetupArchive(createArchive)
			assert.Error(t, err)
			assert.Nil(t, archive)
		}
	}
}
