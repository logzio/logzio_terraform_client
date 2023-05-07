package grafana_folders_test

import (
	"github.com/logzio/logzio_terraform_client/grafana_folders"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"net/http"
	"net/http/httptest"
	"os"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func fixture(path string) string {
	b, err := os.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func setupGrafanaFolderTest() (*grafana_folders.GrafanaFolderClient, error, func()) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, _ := grafana_folders.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}

func setupGrafanaFolderIntegrationTest() (*grafana_folders.GrafanaFolderClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}

	underTest, err := grafana_folders.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}

func getCreateOrUpdateGrafanaFolder() grafana_folders.CreateUpdateFolder {
	return grafana_folders.CreateUpdateFolder{
		Uid:   "client_test",
		Title: "client_test",
	}
}
