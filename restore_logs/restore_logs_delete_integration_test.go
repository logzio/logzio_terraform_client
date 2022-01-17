package restore_logs_test

import (
	"github.com/logzio/logzio_terraform_client/restore_logs"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationRestoreLogs_DeleteRestoreOperation(t *testing.T) {
	underTest, deleteArchive, err := setupRestoreLogsIntegrationTest(withArchive)
	defer test_utils.TestDoneTimeBuffer()
	defer deleteArchive()
	if assert.NoError(t, err) {
		initiateRestore := getInitiateRestoreOperationIntegrationTest()
		restore, err := underTest.InitiateRestoreOperation(initiateRestore)
		if assert.NoError(t, err) && assert.NotNil(t, restore) {
			time.Sleep(2 * time.Second)
			defer func() {
				deleted, err := underTest.DeleteRestoreOperation(restore.Id)
				assert.NoError(t, err)
				assert.NotNil(t, deleted)
				assert.Equal(t, restore.Id, deleted.Id)
				assert.Contains(t, []string{restore_logs.RestoreStatusDeleted, restore_logs.RestoreStatusAborted}, deleted.Status)
			}()
		}
	}
}
