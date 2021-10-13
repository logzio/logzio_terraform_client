package archive_logs_test

import (
	"github.com/logzio/logzio_terraform_client/archive_logs"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationArchiveLogs_UpdateArchiveLogsS3KeysToIam(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()

	if assert.NoError(t, err) {
		createArchive, err := getCreateOrUpdateArchiveLogs(archive_logs.StorageTypeS3)
		if assert.NoError(t, err) {
			archive, err := underTest.SetupArchive(createArchive)
			if assert.NoError(t, err) && assert.NotNil(t, archive) {
				defer underTest.DeleteArchiveLogs(archive.Id)
				time.Sleep(2 * time.Second)
				createArchive.AmazonS3StorageSettings.S3IamCredentials, err = getS3IamCredentials()
				if assert.NoError(t, err) {
					createArchive.AmazonS3StorageSettings.S3SecretCredentials = nil
					createArchive.AmazonS3StorageSettings.CredentialsType = archive_logs.CredentialsTypeIam
					updated, err := underTest.UpdateArchiveLogs(archive.Id, createArchive)
					assert.NoError(t, err)
					assert.NotNil(t, updated)
					assert.Equal(t, archive.Id, updated.Id)
					assert.Equal(t, archive_logs.CredentialsTypeIam, updated.Settings.AmazonS3StorageSettings.CredentialsType)
				}
			}
		}
	}
}

func TestIntegrationArchiveLogs_UpdateArchiveLogsS3toBlob(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()

	if assert.NoError(t, err) {
		createArchive, err := getCreateOrUpdateArchiveLogs(archive_logs.StorageTypeS3)
		if assert.NoError(t, err) {
			archive, err := underTest.SetupArchive(createArchive)
			if assert.NoError(t, err) && assert.NotNil(t, archive) {
				defer underTest.DeleteArchiveLogs(archive.Id)
				time.Sleep(2 * time.Second)
				createArchive, err = getCreateOrUpdateArchiveLogs(archive_logs.StorageTypeBlob)
				if assert.NoError(t, err) {
					updated, err := underTest.UpdateArchiveLogs(archive.Id, createArchive)
					assert.NoError(t, err)
					assert.NotNil(t, updated)
					assert.Equal(t, archive.Id, updated.Id)
					assert.Equal(t, archive_logs.StorageTypeBlob, updated.Settings.StorageType)
				}
			}
		}
	}
}

func TestIntegrationArchiveLogs_UpdateArchiveLogsDisable(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()

	if assert.NoError(t, err) {
		createArchive, err := getCreateOrUpdateArchiveLogs(archive_logs.StorageTypeBlob)
		if assert.NoError(t, err) {
			archive, err := underTest.SetupArchive(createArchive)
			if assert.NoError(t, err) && assert.NotNil(t, archive) {
				defer underTest.DeleteArchiveLogs(archive.Id)
				time.Sleep(2 * time.Second)
				*createArchive.Enabled = false
				if assert.NoError(t, err) {
					updated, err := underTest.UpdateArchiveLogs(archive.Id, createArchive)
					assert.NoError(t, err)
					assert.NotNil(t, updated)
					assert.Equal(t, archive.Id, updated.Id)
					assert.False(t, updated.Settings.Enabled)
				}
			}
		}
	}
}

func TestIntegrationArchiveLogs_UpdateArchiveLogsIdNotFound(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()

	if assert.NoError(t, err) {
		createArchive, err := getCreateOrUpdateArchiveLogs(archive_logs.StorageTypeBlob)
		if assert.NoError(t, err) {
			updated, err := underTest.UpdateArchiveLogs(int32(0000), createArchive)
			assert.Error(t, err)
			assert.Nil(t, updated)
		}
	}
}
