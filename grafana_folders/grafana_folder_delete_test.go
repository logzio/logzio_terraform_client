package grafana_folders_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGrafanaFolder_DeleteGrafanaFolder(t *testing.T) {
	underTest, err, teardown := setupGrafanaFolderTest()
	defer teardown()

	if assert.NoError(t, err) {
		grafanaFolderId := "delete-me"

		mux.HandleFunc("/v1/grafana/api/folders/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			assert.Contains(t, r.URL.String(), grafanaFolderId)
		})

		err = underTest.DeleteGrafanaFolder(grafanaFolderId)
		assert.NoError(t, err)
	}
}

func TestGrafanaFolder_DeleteGrafanaFolderInternalServerError(t *testing.T) {
	underTest, err, teardown := setupGrafanaFolderTest()
	defer teardown()

	if assert.NoError(t, err) {
		grafanaFolderId := "delete-me"

		mux.HandleFunc("/v1/grafana/api/folders/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			assert.Contains(t, r.URL.String(), grafanaFolderId)
			w.WriteHeader(http.StatusInternalServerError)
		})

		err = underTest.DeleteGrafanaFolder(grafanaFolderId)
		assert.Error(t, err)
	}
}
