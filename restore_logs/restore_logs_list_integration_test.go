package restore_logs_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationRestoreLogs_ListRestoreOperations(t *testing.T) {
	underTest, _, err := setupRestoreLogsIntegrationTest(withoutArchive)

	if assert.NoError(t, err) {
		archives, err := underTest.ListRestoreOperations()
		assert.NoError(t, err)
		assert.NotNil(t, archives)
	}
}
