package grafana_folders_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationGrafanaFolder_ListGrafanaFolders(t *testing.T) {
	underTest, err := setupGrafanaFolderIntegrationTest()

	if assert.NoError(t, err) {
		folders, err := underTest.ListGrafanaFolders()
		assert.NoError(t, err)
		assert.NotNil(t, folders)
	}
}
