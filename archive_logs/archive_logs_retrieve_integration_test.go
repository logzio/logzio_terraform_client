package archive_logs_test

import (
	"github.com/logzio/logzio_terraform_client/archive_logs"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationArchiveLogs_RetrieveArchive(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createArchive, err := test_utils.GetCreateOrUpdateArchiveLogs(archive_logs.StorageTypeS3)
		if assert.NoError(t, err) {
			archive, err := underTest.SetupArchive(createArchive)
			if assert.NoError(t, err) && assert.NotNil(t, archive) {
				defer underTest.DeleteArchiveLogs(archive.Id)
				time.Sleep(2 * time.Second)
				getArchive, err := underTest.RetrieveArchiveLogsSetting(archive.Id)
				assert.NoError(t, err)
				assert.NotNil(t, getArchive)
				assert.Equal(t, archive.Id, getArchive.Id)
				assert.Equal(t, archive.Settings, getArchive.Settings)
			}
		}
	}
}

func TestIntegrationArchiveLogs_RetrieveArchiveIdNotFound(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		archive, err := underTest.RetrieveArchiveLogsSetting(int32(1234))
		assert.Error(t, err)
		assert.Nil(t, archive)
	}
}
