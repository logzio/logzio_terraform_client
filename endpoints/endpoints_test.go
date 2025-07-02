package endpoints_test

import (
	"github.com/logzio/logzio_terraform_client/endpoints"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

const (
	testsUrl       = "https://jsonplaceholder.typicode.com/todos/1"
	testsUrlUpdate = "https://jsonplaceholder.typicode.com/todos/2"
)

func fixture(path string) string {
	b, err := os.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func setupEndpointsTest() (*endpoints.EndpointsClient, error, func()) {
	apiToken := "SOME_API_TOKEN"

	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	underTest, _ := endpoints.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}

func setupEndpointsIntegrationTest() (*endpoints.EndpointsClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := endpoints.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, nil
}

func TestNewWithEmptyBaseUrl(t *testing.T) {
	_, err := endpoints.New("any-api-token", "")
	if err == nil {
		t.Fatal("Expected error when baseUrl is empty")
	}
	if err.Error() != "Base URL not defined" {
		t.Fatalf("The expected error message to be '%s' but was '%s'",
			"Base URL not defined", err.Error())
	}
}

func TestNewWithEmptyApiToken(t *testing.T) {
	_, err := endpoints.New("", "any-base-url")

	if err == nil {
		t.Fatal("Expected error when API token is empty")
	}
	if err.Error() != "API token not defined" {
		t.Fatalf("The expected error message to be '%s' but was '%s'",
			"API token not defined", err.Error())
	}
}

func GetCreateOrUpdateEndpoint() endpoints.CreateOrUpdateEndpoint {
	return endpoints.CreateOrUpdateEndpoint{
		Title:       "tf_test",
		Description: "this is a description",
	}
}
