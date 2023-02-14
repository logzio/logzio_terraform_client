package archive_logs_test

import (
	"github.com/logzio/logzio_terraform_client/archive_logs"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationArchiveLogs_DeleteArchive(t *testing.T) {
	underTest, err := setupArchiveLogsIntegrationTest()

	if assert.NoError(t, err) {
		createArchive, err := test_utils.GetCreateOrUpdateArchiveLogs(archive_logs.StorageTypeS3)
		if assert.NoError(t, err) {
			archive, err := underTest.SetupArchive(createArchive)
			if assert.NoError(t, err) && assert.NotNil(t, archive) {
				time.Sleep(2 * time.Second)
				defer func() {
					err = underTest.DeleteArchiveLogs(archive.Id)
					assert.NoError(t, err)
				}()
			}
		}
	}
}
