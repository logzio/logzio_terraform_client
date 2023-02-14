package restore_logs_test

import (
	"github.com/logzio/logzio_terraform_client/restore_logs"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationRestoreLogs_InitiateRestore(t *testing.T) {
	underTest, deleteArchive, err := setupRestoreLogsIntegrationTest(withArchive)
	defer deleteArchive()
	if assert.NoError(t, err) {
		initiateRestore := getInitiateRestoreOperationIntegrationTest()
		restore, err := underTest.InitiateRestoreOperation(initiateRestore)
		if assert.NoError(t, err) && assert.NotNil(t, restore) {
			time.Sleep(2 * time.Second)
			defer underTest.DeleteRestoreOperation(restore.Id)
			assert.NotEmpty(t, restore.Id)
			assert.Contains(t, []string{restore_logs.RestoreStatusInProgress, restore_logs.RestoreStatusActive}, restore.Status)
		}
	}
}

func TestIntegrationRestoreLogs_InitiateRestoreEmptyName(t *testing.T) {
	underTest, _, err := setupRestoreLogsIntegrationTest(withoutArchive)
	if assert.NoError(t, err) {
		initiateRestore := getInitiateRestoreOperationIntegrationTest()
		initiateRestore.AccountName = ""
		restore, err := underTest.InitiateRestoreOperation(initiateRestore)
		assert.Error(t, err)
		assert.Nil(t, restore)
	}
}

func TestIntegrationRestoreLogs_InitiateRestoreEmptyStartTime(t *testing.T) {
	underTest, _, err := setupRestoreLogsIntegrationTest(withoutArchive)
	if assert.NoError(t, err) {
		initiateRestore := getInitiateRestoreOperationIntegrationTest()
		initiateRestore.StartTime = 0
		restore, err := underTest.InitiateRestoreOperation(initiateRestore)
		assert.Error(t, err)
		assert.Nil(t, restore)
	}
}

func TestIntegrationRestoreLogs_InitiateRestoreEmptyEndTime(t *testing.T) {
	underTest, _, err := setupRestoreLogsIntegrationTest(withoutArchive)
	if assert.NoError(t, err) {
		initiateRestore := getInitiateRestoreOperationIntegrationTest()
		initiateRestore.EndTime = 0
		restore, err := underTest.InitiateRestoreOperation(initiateRestore)
		assert.Error(t, err)
		assert.Nil(t, restore)
	}
}
