package drop_metrics_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"

	"github.com/logzio/logzio_terraform_client/drop_metrics"
)

var (
	mux    *http.ServeMux
	client *drop_metrics.DropMetricsClient
	server *httptest.Server
)

func setupDropMetricsTest() (*drop_metrics.DropMetricsClient, *http.ServeMux, func()) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client, _ = drop_metrics.New("test-token", server.URL)

	return client, mux, server.Close
}

func fixture(filename string) string {
	content, err := os.ReadFile(filepath.Join("testdata", "fixtures", filename))
	if err != nil {
		return fmt.Sprintf("fixture file not found: %s", filename)
	}
	return string(content)
}
