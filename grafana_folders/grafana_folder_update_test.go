package grafana_folders_test

import (
	"encoding/json"
	"github.com/logzio/logzio_terraform_client/grafana_folders"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestS3Fetcher_UpdateFetcher(t *testing.T) {
	underTest, err, teardown := setupGrafanaFolderTest()
	assert.NoError(t, err)
	defer teardown()

	folderId := "client_test"

	mux.HandleFunc("/v1/grafana/api/folders/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), folderId)
		jsonBytes, _ := io.ReadAll(r.Body)
		var target grafana_folders.CreateUpdateFolder
		err = json.Unmarshal(jsonBytes, &target)
		assert.NoError(t, err)
		assert.NotNil(t, target)
		assert.NotEmpty(t, target.Title)
		assert.Equal(t, folderId, target.Uid)
		assert.True(t, target.Overwrite)
	})

	updateFolder := getCreateOrUpdateGrafanaFolder()
	err = underTest.UpdateGrafanaFolder(updateFolder)
	assert.NoError(t, err)
}
