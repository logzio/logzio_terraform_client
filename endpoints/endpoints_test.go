package endpoints_test

import (
	"github.com/jonboydell/logzio_client/endpoints"
	"github.com/jonboydell/logzio_client/test_utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
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
