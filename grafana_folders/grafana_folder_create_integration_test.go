package grafana_folders_test

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationGrafanaFolder_CreateGrafanaFolder(t *testing.T) {
	underTest, err := setupGrafanaFolderIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createGrafanaFolder := getCreateOrUpdateGrafanaFolder()
		createGrafanaFolder.Title = fmt.Sprintf("%s_%s", createGrafanaFolder.Title, "create")
		grafanaFolder, err := underTest.CreateGrafanaFolder(createGrafanaFolder)
		if assert.NoError(t, err) && assert.NotNil(t, grafanaFolder) {
			time.Sleep(4 * time.Second)
			defer underTest.DeleteGrafanaFolder(grafanaFolder.Uid)
			assert.Equal(t, createGrafanaFolder.Uid, grafanaFolder.Uid)
			assert.Equal(t, createGrafanaFolder.Title, grafanaFolder.Title)
			assert.Equal(t, int64(1), grafanaFolder.Version)
		}
	}
}

func TestIntegrationGrafanaFolder_CreateGrafanaFolderNoTitle(t *testing.T) {
	underTest, err := setupGrafanaFolderIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createGrafanaFolder := getCreateOrUpdateGrafanaFolder()
		createGrafanaFolder.Title = ""
		grafanaFolder, err := underTest.CreateGrafanaFolder(createGrafanaFolder)
		assert.Error(t, err)
		assert.Nil(t, grafanaFolder)
	}
}

func TestIntegrationGrafanaFolder_CreateGrafanaFolderInvalidTitle(t *testing.T) {
	underTest, err := setupGrafanaFolderIntegrationTest()
	defer test_utils.TestDoneTimeBuffer()

	if assert.NoError(t, err) {
		createGrafanaFolder := getCreateOrUpdateGrafanaFolder()

		// test `/` naming limitation
		createGrafanaFolder.Title = "client/test/title"
		grafanaFolder, err := underTest.CreateGrafanaFolder(createGrafanaFolder)
		assert.Error(t, err)
		assert.Nil(t, grafanaFolder)

		// test `\` naming limitation
		createGrafanaFolder.Title = "client\\test\\title"
		grafanaFolder, err = underTest.CreateGrafanaFolder(createGrafanaFolder)
		assert.Error(t, err)
		assert.Nil(t, grafanaFolder)
	}
}
