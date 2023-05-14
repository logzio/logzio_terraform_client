package grafana_folders_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationGrafanaFolder_UpdateGrafanaFolder(t *testing.T) {
	underTest, err := setupGrafanaFolderIntegrationTest()

	if assert.NoError(t, err) {
		request := getCreateOrUpdateGrafanaFolder()
		createFolder, err := underTest.CreateGrafanaFolder(request)
		if assert.NoError(t, err) && assert.NotNil(t, createFolder) && assert.NotEmpty(t, createFolder.Uid) {
			assert.Equal(t, int64(1), createFolder.Version)
			defer underTest.DeleteGrafanaFolder(createFolder.Uid)
			time.Sleep(time.Second * 2)
			newTitle := fmt.Sprintf("%s_%s", request.Title, "update")
			request.Title = newTitle
			request.Overwrite = true
			err = underTest.UpdateGrafanaFolder(request)
			assert.NoError(t, err)
			// verify that the update was made
			time.Sleep(time.Second * 2)
			getFolder, err := underTest.GetGrafanaFolder(createFolder.Uid)
			assert.NoError(t, err)
			assert.Equal(t, newTitle, getFolder.Title)
			assert.Equal(t, int64(2), getFolder.Version)
		}
	}
}

func TestIntegrationGrafanaFolder_UpdateGrafanaFolderIdNotFound(t *testing.T) {
	underTest, err := setupGrafanaFolderIntegrationTest()

	if assert.NoError(t, err) {
		request := getCreateOrUpdateGrafanaFolder()
		request.Uid = "not-exist"
		err = underTest.UpdateGrafanaFolder(request)
		assert.Error(t, err)
	}
}
