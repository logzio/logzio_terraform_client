package grafana_folders_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGrafanaFolder_GetGrafanaFolder(t *testing.T) {
	underTest, err, teardown := setupGrafanaFolderTest()
	assert.NoError(t, err)
	defer teardown()

	folderUid := "client_test"

	mux.HandleFunc("/v1/grafana/api/folders/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), folderUid)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("create_grafana_folder_res.json"))
	})

	folder, err := underTest.GetGrafanaFolder(folderUid)
	assert.NoError(t, err)
	assert.NotNil(t, folder)
	assert.Equal(t, folderUid, folder.Uid)
	assert.Equal(t, "client_test_title", folder.Title)
	assert.Equal(t, "/grafana-app/dashboards/f/client_test/client_test", folder.Url)
	assert.Equal(t, int64(1), folder.Version)
	assert.Equal(t, int64(123456), folder.Id)
}

func TestGrafanaFolder_GetGrafanaFolderInternalError(t *testing.T) {
	underTest, err, teardown := setupGrafanaFolderTest()
	assert.NoError(t, err)
	defer teardown()

	folderUid := "some-id"

	mux.HandleFunc("/v1/grafana/api/folders/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), folderUid)
		w.WriteHeader(http.StatusInternalServerError)
	})

	folder, err := underTest.GetGrafanaFolder(folderUid)
	assert.Error(t, err)
	assert.Nil(t, folder)
}
