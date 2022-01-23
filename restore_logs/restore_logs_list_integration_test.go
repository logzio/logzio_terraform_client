package restore_logs_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationRestoreLogs_ListRestoreOperations(t *testing.T) {
	underTest, _, err := setupRestoreLogsIntegrationTest(withoutArchive)
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		restores, err := underTest.ListRestoreOperations()
		assert.NoError(t, err)
		assert.NotNil(t, restores)
	}
}
