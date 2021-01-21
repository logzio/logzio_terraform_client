package shared_tokens_test

import (
	"github.com/logzio/logzio_terraform_client/shared_tokens"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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

func setupSharedTokensTest() (*shared_tokens.SharedTokenClient, error, func()) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, _ := shared_tokens.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}

func setupSharedTokensIntegrationTest() (*shared_tokens.SharedTokenClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}

	underTest, err := shared_tokens.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}
