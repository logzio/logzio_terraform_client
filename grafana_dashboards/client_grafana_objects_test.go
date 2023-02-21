package grafana_dashboards_test

import (
	"github.com/logzio/logzio_terraform_client/grafana_dashboards"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
)

const dashboardsApiBasePath = "/v1/grafana/api/dashboards"

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func setupGrafanaObjectsTest() (*grafana_dashboards.GrafanaObjectsClient, func(), error) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, err := grafana_dashboards.New(apiToken, server.URL)

	return underTest, func() {
		server.Close()
	}, err
}

func setupGrafanaObjectsIntegrationTest() (*grafana_dashboards.GrafanaObjectsClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := grafana_dashboards.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}

func fixture(path string) string {
	b, err := os.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func getCreateUpdateDashboard() grafana_dashboards.CreateUpdatePayload {
	dashboard := map[string]interface{}{
		"title":  "dashboard_test",
		"tags":   []string{"some", "tags"},
		"panels": make([]interface{}, 0),
	}

	return grafana_dashboards.CreateUpdatePayload{
		Dashboard: dashboard,
		FolderId:  1,
		Message:   "some message",
		Overwrite: true,
	}
}

func getCreateDashboardIntegrationTests() (grafana_dashboards.CreateUpdatePayload, error) {
	createUpdate := getCreateUpdateDashboard()
	folderId, err := test_utils.GetMetricsFolderId()
	if err != nil {
		return createUpdate, err
	}

	createUpdate.FolderId, err = strconv.Atoi(folderId)
	return createUpdate, err
}
