package archive_logs_test

import (
	"github.com/logzio/logzio_terraform_client/archive_logs"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationArchiveLogs_ListArchives(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		archives, err := underTest.ListArchiveLog()
		assert.NoError(t, err)
		assert.NotNil(t, archives)
	}
}

func TestIntegrationArchiveLogs_ListArchivesAtLeastOneArchive(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createArchive, err := test_utils.GetCreateOrUpdateArchiveLogs(archive_logs.StorageTypeS3)
		if assert.NoError(t, err) {
			archive, err := underTest.SetupArchive(createArchive)
			if assert.NoError(t, err) && assert.NotNil(t, archive) {
				defer underTest.DeleteArchiveLogs(archive.Id)
				time.Sleep(2 * time.Second)
				archives, err := underTest.ListArchiveLog()
				assert.NoError(t, err)
				assert.NotNil(t, archives)
				assert.True(t, len(archives) > 0)
			}

		}
	}
}
