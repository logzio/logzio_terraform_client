package sub_accounts_test

import (
	"github.com/jonboydell/logzio_client/sub_accounts"
	"github.com/jonboydell/logzio_client/test_utils"
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

func setupSubAccountsTest() (*sub_accounts.SubAccountClient, error, func()) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, _ := sub_accounts.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}

func setupSubAccountsIntegrationTest() (*sub_accounts.SubAccountClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}

	underTest, err := sub_accounts.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}
