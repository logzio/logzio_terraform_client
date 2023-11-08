package grafana_contact_points_test

import (
	"github.com/logzio/logzio_terraform_client/grafana_contact_points"
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

func setupGrafanaContactPointIntegrationTest() (*grafana_contact_points.GrafanaContactPointClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}

	underTest, err := grafana_contact_points.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}

func getGrafanaContactPointObject() grafana_contact_points.GrafanaContactPoint {
	return grafana_contact_points.GrafanaContactPoint{
		Name: "tf-client-test",
		Type: "email",
		Settings: map[string]interface{}{
			"addresses":   "example1@example.com;example2@example.com",
			"singleEmail": false,
		},
		DisableResolveMessage: false,
	}
}
