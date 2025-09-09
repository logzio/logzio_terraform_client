package unified_projects_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/logzio/logzio_terraform_client/unified_projects"
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

func setupUnifiedProjectsTest() (*unified_projects.ProjectsClient, error, func()) {
	apiToken := "SOME_API_TOKEN"

	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	underTest, _ := unified_projects.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}

func TestNewWithEmptyBaseUrl(t *testing.T) {
	_, err := unified_projects.New("any-api-token", "")
	if err == nil {
		t.Fatal("Expected error when baseUrl is empty")
	}
	if err.Error() != "Base URL not defined" {
		t.Fatalf("The expected error message to be '%s' but was '%s'",
			"Base URL not defined", err.Error())
	}
}

func TestNewWithEmptyApiToken(t *testing.T) {
	_, err := unified_projects.New("", "any-base-url")

	if err == nil {
		t.Fatal("Expected error when API token is empty")
	}
	if err.Error() != "API token not defined" {
		t.Fatalf("The expected error message to be '%s' but was '%s'",
			"API token not defined", err.Error())
	}
}

func setupUnifiedProjectsIntegrationTest() (*unified_projects.ProjectsClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := unified_projects.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}
