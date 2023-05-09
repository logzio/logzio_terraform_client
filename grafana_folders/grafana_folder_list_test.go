package grafana_folders_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGrafanaFolder_ListGrafanaFolders(t *testing.T) {
	underTest, err, teardown := setupGrafanaFolderTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/grafana/api/folders", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("list_grafana_folders.json"))
	})

	folders, err := underTest.ListGrafanaFolders()
	assert.NoError(t, err)
	assert.NotNil(t, folders)
	assert.Equal(t, 2, len(folders))
}

func TestGrafanaFolder_ListGrafanaFoldersInternalServerError(t *testing.T) {
	underTest, err, teardown := setupGrafanaFolderTest()
	assert.NoError(t, err)
	defer teardown()

	mux.HandleFunc("/v1/grafana/api/folders", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
	})

	folders, err := underTest.ListGrafanaFolders()
	assert.Error(t, err)
	assert.Nil(t, folders)
}
