package archive_logs_test

// TODO
//import (
//	"github.com/logzio/logzio_terraform_client/archive_logs"
//	"github.com/stretchr/testify/assert"
//	"testing"
//	"time"
//)
//
//func TestIntegrationArchiveLogs_RetrieveArchiveS3(t *testing.T) {
//	underTest, err := setupArchiveLogsIntegrationTest()
//
//	if assert.NoError(t, err) {
//		createArchive, err := getCreateOrUpdateArchiveLogs(archive_logs.StorageTypeS3)
//		if assert.NoError(t, err) {
//			archive, err := underTest.SetupArchive(createArchive)
//			if assert.NoError(t, err) && assert.NotNil(t, archive) {
//				defer underTest.DeleteArchiveLogs(archive.Id)
//				time.Sleep(2 * time.Second)
//				getArchive, err := underTest.RetrieveArchiveLogsSetting(archive.Id)
//				assert.NoError(t, err)
//				assert.NotNil(t, getArchive)
//				assert.Equal(t, archive.Id, getArchive.Id)
//				assert.Equal(t, archive.Settings, getArchive.Settings)
//			}
//		}
//
//	}
//}
