package grafana_objects_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/logzio/logzio_terraform_client/grafana_objects"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func setupGrafanaObjectsTest() (*grafana_objects.GrafanaObjectsClient, error, func()) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, _ := grafana_objects.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}
