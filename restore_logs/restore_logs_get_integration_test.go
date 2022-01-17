package restore_logs_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationRestoreLogs_GetRestoreOperation(t *testing.T) {
	underTest, deleteArchive, err := setupRestoreLogsIntegrationTest(withArchive)
	defer test_utils.TestDoneTimeBuffer()
	defer deleteArchive()
	if assert.NoError(t, err) {
		initiateRestore := getInitiateRestoreOperationIntegrationTest()
		restore, err := underTest.InitiateRestoreOperation(initiateRestore)
		if assert.NoError(t, err) && assert.NotNil(t, restore) {
			defer underTest.DeleteRestoreOperation(restore.Id)
			time.Sleep(2 * time.Second)
			getRestore, err := underTest.GetRestoreOperation(restore.Id)
			assert.NoError(t, err)
			assert.NotNil(t, getRestore)
			assert.Equal(t, restore.Id, getRestore.Id)
		}
	}
}

func TestIntegrationRestoreLogs_GetRestoreOperationInvalidId(t *testing.T) {
	underTest, _, err := setupRestoreLogsIntegrationTest(withoutArchive)

	if assert.NoError(t, err) {
		getRestore, err := underTest.GetRestoreOperation(int32(1))
		assert.Error(t, err)
		assert.Nil(t, getRestore)
	}
}
