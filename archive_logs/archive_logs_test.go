package archive_logs_test

import (
	"github.com/logzio/logzio_terraform_client/archive_logs"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

const archiveApiBasePath = "/v2/archive/settings"

func setupArchiveLogsTest() (*archive_logs.ArchiveLogsClient, func(), error) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, _ := archive_logs.New(apiToken, server.URL)

	return underTest, func() {
		server.Close()
	}, nil
}

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func setupArchiveLogsIntegrationTest() (*archive_logs.ArchiveLogsClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}

	underTest, err := archive_logs.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}
