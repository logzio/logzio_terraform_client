package grafana_datasources_test

import (
	"github.com/logzio/logzio_terraform_client/grafana_datasources"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"net/http"
	"net/http/httptest"
	"os"
)

const (
	envMetricsAccountName = "METRICS_ACCOUNT_NAME"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func setupGrafanaDatasourceIntegrationTest() (*grafana_datasources.GrafanaDatasourceClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := grafana_datasources.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}

func setupGrafanaDatasourceTest() (*grafana_datasources.GrafanaDatasourceClient, func(), error) {
	apiToken := "SOME_API_TOKEN"

	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	underTest, _ := grafana_datasources.New(apiToken, server.URL)

	return underTest, func() {
		server.Close()
	}, nil
}

func fixture(path string) string {
	b, err := os.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
