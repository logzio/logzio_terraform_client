package grafana_folders_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationGrafanaFolder_GetGrafanaFolder(t *testing.T) {
	underTest, err := setupGrafanaFolderIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createGrafanaFolder := getCreateOrUpdateGrafanaFolder()
		grafanaFolder, err := underTest.CreateGrafanaFolder(createGrafanaFolder)
		if assert.NoError(t, err) && assert.NotNil(t, grafanaFolder) && assert.NotEmpty(t, grafanaFolder.Uid) {
			defer underTest.DeleteGrafanaFolder(grafanaFolder.Uid)
			time.Sleep(4 * time.Second)
			getFolder, err := underTest.GetGrafanaFolder(grafanaFolder.Uid)
			assert.NoError(t, err)
			assert.NotNil(t, getFolder)
			assert.Equal(t, grafanaFolder, getFolder)
		}
	}
}

func TestIntegrationGrafanaFolder_GetGrafanaFolderUidNotExists(t *testing.T) {
	underTest, err := setupGrafanaFolderIntegrationTest()

	if assert.NoError(t, err) {
		folder, err := underTest.GetGrafanaFolder("not-exist")
		assert.Error(t, err)
		assert.Nil(t, folder)
		assert.Contains(t, err.Error(), "failed with missing grafana folder")
	}
}
