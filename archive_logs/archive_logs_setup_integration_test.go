package archive_logs_test

import (
	"github.com/logzio/logzio_terraform_client/archive_logs"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationArchiveLogs_SetupArchiveS3Keys(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()
	if assert.NoError(t, err) {
		createArchiveLogs, err := getCreateOrUpdateArchiveLogs(archive_logs.StorageTypeS3)
		if assert.NoError(t, err) {
			archiveLogs, err := underTest.SetupArchive(createArchiveLogs)
			if assert.NoError(t, err) && assert.NotNil(t, archiveLogs) {
				time.Sleep(2 * time.Second)
				defer underTest.DeleteArchiveLogs(archiveLogs.Id)
				assert.NotEmpty(t, archiveLogs.Id)
				assert.NotEmpty(t, archiveLogs.Settings)
			}
		}
	}
}

func TestIntegrationArchiveLogs_SetupArchiveS3Iam(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()
	if assert.NoError(t, err) {
		createArchiveLogs, err := getCreateOrUpdateArchiveLogs(archive_logs.StorageTypeS3)
		createArchiveLogs.AmazonS3StorageSettings.CredentialsType = archive_logs.CredentialsTypeIam
		createArchiveLogs.AmazonS3StorageSettings.S3SecretCredentials = nil
		if assert.NoError(t, err) {
			iam, err := getS3IamCredentials()
			if assert.NoError(t, err) {
				createArchiveLogs.AmazonS3StorageSettings.S3IamCredentials = iam
				if assert.NoError(t, err) {
					archiveLogs, err := underTest.SetupArchive(createArchiveLogs)
					if assert.NoError(t, err) && assert.NotNil(t, archiveLogs) {
						time.Sleep(4 * time.Second)
						defer underTest.DeleteArchiveLogs(archiveLogs.Id)
						assert.NotEmpty(t, archiveLogs.Id)
						assert.NotEmpty(t, archiveLogs.Settings)
					}
				}
			}
		}
	}
}

func TestIntegrationArchiveLogs_SetupArchiveBlob(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()
	if assert.NoError(t, err) {
		createArchiveLogs, err := getCreateOrUpdateArchiveLogs(archive_logs.StorageTypeBlob)
		if assert.NoError(t, err) {
			archiveLogs, err := underTest.SetupArchive(createArchiveLogs)
			if assert.NoError(t, err) && assert.NotNil(t, archiveLogs) {
				time.Sleep(4 * time.Second)
				defer underTest.DeleteArchiveLogs(archiveLogs.Id)
				assert.NotEmpty(t, archiveLogs.Id)
				assert.NotEmpty(t, archiveLogs.Settings)
			}
		}
	}
}

func TestIntegrationArchiveLogs_SetupArchiveInvalidStorageType(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()

	if assert.NoError(t, err) {
		createArchiveLogs, _ := getCreateOrUpdateArchiveLogs(archive_logs.StorageTypeBlob)
		createArchiveLogs.StorageType = "invalid type"
		archive, err := underTest.SetupArchive(createArchiveLogs)

		assert.Error(t, err)
		assert.Nil(t, archive)
	}
}

func TestIntegrationArchiveLogs_SetupArchiveS3InvalidAccessKey(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()

	if assert.NoError(t, err) {
		createArchiveLogs, _ := getCreateOrUpdateArchiveLogs(archive_logs.StorageTypeS3)
		createArchiveLogs.AmazonS3StorageSettings.S3SecretCredentials.AccessKey = "invalid access key"
		archive, err := underTest.SetupArchive(createArchiveLogs)

		assert.Error(t, err)
		assert.Nil(t, archive)
	}
}

func TestIntegrationArchiveLogs_SetupArchiveS3InvalidSecretKey(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()

	if assert.NoError(t, err) {
		createArchiveLogs, _ := getCreateOrUpdateArchiveLogs(archive_logs.StorageTypeS3)
		createArchiveLogs.AmazonS3StorageSettings.S3SecretCredentials.SecretKey = "invalid secret key"
		archive, err := underTest.SetupArchive(createArchiveLogs)

		assert.Error(t, err)
		assert.Nil(t, archive)
	}
}

func TestIntegrationArchiveLogs_SetupArchiveBlobInvalidTenantId(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()

	if assert.NoError(t, err) {
		createArchiveLogs, _ := getCreateOrUpdateArchiveLogs(archive_logs.StorageTypeBlob)
		createArchiveLogs.AzureBlobStorageSettings.TenantId = "invalid tenant id"
		archive, err := underTest.SetupArchive(createArchiveLogs)

		assert.Error(t, err)
		assert.Nil(t, archive)
	}
}

func TestIntegrationArchiveLogs_SetupArchiveBlobInvalidClientId(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()

	if assert.NoError(t, err) {
		createArchiveLogs, _ := getCreateOrUpdateArchiveLogs(archive_logs.StorageTypeBlob)
		createArchiveLogs.AzureBlobStorageSettings.ClientId = "invalid client id"
		archive, err := underTest.SetupArchive(createArchiveLogs)

		assert.Error(t, err)
		assert.Nil(t, archive)
	}
}

func TestIntegrationArchiveLogs_SetupArchiveBlobInvalidClientSecret(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()

	if assert.NoError(t, err) {
		createArchiveLogs, _ := getCreateOrUpdateArchiveLogs(archive_logs.StorageTypeBlob)
		createArchiveLogs.AzureBlobStorageSettings.ClientSecret = "invalid client secret"
		archive, err := underTest.SetupArchive(createArchiveLogs)

		assert.Error(t, err)
		assert.Nil(t, archive)
	}
}

func TestIntegrationArchiveLogs_SetupArchiveBlobInvalidAccountName(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()

	if assert.NoError(t, err) {
		createArchiveLogs, _ := getCreateOrUpdateArchiveLogs(archive_logs.StorageTypeBlob)
		createArchiveLogs.AzureBlobStorageSettings.AccountName = "invalid account name"
		archive, err := underTest.SetupArchive(createArchiveLogs)

		assert.Error(t, err)
		assert.Nil(t, archive)
	}
}

func TestIntegrationArchiveLogs_SetupArchiveBlobInvalidContainerName(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()

	if assert.NoError(t, err) {
		createArchiveLogs, _ := getCreateOrUpdateArchiveLogs(archive_logs.StorageTypeBlob)
		createArchiveLogs.AzureBlobStorageSettings.ContainerName = "invalid container name"
		archive, err := underTest.SetupArchive(createArchiveLogs)

		assert.Error(t, err)
		assert.Nil(t, archive)
	}
}
