package grafana_folders_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/grafana_folders"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestGrafanaFolder_CreateGrafanaFolder(t *testing.T) {
	underTest, err, teardown := setupGrafanaFolderTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/grafana/api/folders", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target grafana_folders.CreateUpdateFolder
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Title)
			assert.NotEmpty(t, target.Uid)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("create_grafana_folder_res.json"))
		})

		createGrafanaFolder := getCreateOrUpdateGrafanaFolder()
		grafanaFolder, err := underTest.CreateGrafanaFolder(createGrafanaFolder)
		assert.NoError(t, err)
		assert.NotNil(t, grafanaFolder)
		assert.Equal(t, int64(123456), grafanaFolder.Id)
		assert.Equal(t, createGrafanaFolder.Uid, grafanaFolder.Uid)
		assert.Equal(t, createGrafanaFolder.Title, grafanaFolder.Title)
		assert.Equal(t, "/grafana-app/dashboards/f/client_test/client_test", grafanaFolder.Url)
		assert.Equal(t, int64(1), grafanaFolder.Version)
	}
}

func TestGrafanaFolder_CreateGrafanaFolderInternalServerError(t *testing.T) {
	underTest, err, teardown := setupGrafanaFolderTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/grafana/api/folders", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target grafana_folders.CreateUpdateFolder
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Title)
			assert.NotEmpty(t, target.Uid)
			w.WriteHeader(http.StatusInternalServerError)
		})

		createGrafanaFolder := getCreateOrUpdateGrafanaFolder()
		grafanaFolder, err := underTest.CreateGrafanaFolder(createGrafanaFolder)
		assert.Error(t, err)
		assert.Nil(t, grafanaFolder)
	}
}

func TestGrafanaFolder_CreateGrafanaFolderInvalidFolderNameError(t *testing.T) {
	underTest, err, teardown := setupGrafanaFolderTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/grafana/api/folders", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := io.ReadAll(r.Body)
			var target grafana_folders.CreateUpdateFolder
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Title)
			assert.NotEmpty(t, target.Uid)
			w.Header().Set("Content-Type", "application/json")
		})
		// test `/` naming limitation
		createGrafanaFolder := getCreateOrUpdateGrafanaFolder()
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
