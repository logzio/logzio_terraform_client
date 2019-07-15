package endpoints_test

import (
	"github.com/jonboydell/logzio_client/client"
	"github.com/jonboydell/logzio_client/endpoints"
	"github.com/jonboydell/logzio_client/test_utils"
	"net/http"
	"net/http/httptest"
)

var (
	mux *http.ServeMux
	server *httptest.Server
)

func setupEndpointsTest() (*endpoints.EndpointsClient, error, func()) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err, nil
	}

	underTest, err := endpoints.New(apiToken)

	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	underTest.BaseUrl = server.URL
	underTest.Client.BaseUrl = server.URL

	return underTest, nil, func() {
		server.Close()
	}
}

func setupEndpointsIntegrationTest() (*endpoints.EndpointsClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := endpoints.New(apiToken)
	underTest.BaseUrl = client.GetLogzIoBaseUrl()
	return underTest, nil
}