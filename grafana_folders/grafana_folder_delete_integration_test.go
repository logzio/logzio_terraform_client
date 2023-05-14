package grafana_folders_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationGrafanaFolder_DeleteGrafanaFolder(t *testing.T) {
	underTest, err := setupGrafanaFolderIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()
	if assert.NoError(t, err) {
		createGrafanaFolder := getCreateOrUpdateGrafanaFolder()
		grafanaFolder, err := underTest.CreateGrafanaFolder(createGrafanaFolder)
		if assert.NoError(t, err) && assert.NotNil(t, grafanaFolder) && assert.NotEmpty(t, grafanaFolder.Uid) {
			time.Sleep(2 * time.Second)
			defer func() {
				err = underTest.DeleteGrafanaFolder(grafanaFolder.Uid)
				assert.NoError(t, err)
			}()
		}
	}
}

func TestIntegrationGrafanaFolder_DeleteGrafanaFolderNotExists(t *testing.T) {
	underTest, err := setupGrafanaFolderIntegrationTest()
	if assert.NoError(t, err) {
		err = underTest.DeleteGrafanaFolder("not-exists")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed with missing grafana folder")
	}
}
