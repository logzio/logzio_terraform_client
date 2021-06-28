package log_shipping_tokens_test

import (
	"github.com/logzio/logzio_terraform_client/log_shipping_tokens"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func setupLogShippingTokensIntegrationTest() (*log_shipping_tokens.LogShippingTokensClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := log_shipping_tokens.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, nil
}

func getCreateLogShippingToken() log_shipping_tokens.CreateLogShippingToken {
	return log_shipping_tokens.CreateLogShippingToken{
		Name:    "client_integration_test",
	}
}

func setupLogShippingTokenTest() (*log_shipping_tokens.LogShippingTokensClient, error, func()) {
	apiToken := "SOME_API_TOKEN"

	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	underTest, _ := log_shipping_tokens.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}