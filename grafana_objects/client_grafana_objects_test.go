package grafana_objects_test

import (
	"github.com/logzio/logzio_terraform_client/grafana_objects"
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

func setupGrafanaObjectsTest() (*grafana_objects.GrafanaObjectsClient, func(), error) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, err := grafana_objects.New(apiToken, server.URL)

	return underTest, func() {
		server.Close()
	}, err
}

func setupGrafanaObjectsIntegrationTest() (*grafana_objects.GrafanaObjectsClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := grafana_objects.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}

func fixture(path string) string {
	b, err := os.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func getCreateUpdateDashboard() grafana_objects.CreateUpdatePayload {
	dashboard := grafana_objects.DashboardObject{
		Title:  "dashboard_test",
		Tags:   []string{"some", "tags"},
		Panels: make([]interface{}, 0),
	}

	return grafana_objects.CreateUpdatePayload{
		Dashboard: dashboard,
		FolderId:  1,
		Message:   "some message",
		Overwrite: true,
	}
}

func getCreateDashboardIntegrationTests() (grafana_objects.CreateUpdatePayload, error) {
	createUpdate := getCreateUpdateDashboard()
	folderId, err := test_utils.GetMetricsFolderId()
	if err != nil {
		return createUpdate, err
	}

	createUpdate.FolderId, err = strconv.Atoi(folderId)
	return createUpdate, err
}
