package users_test

import (
	"github.com/jonboydell/logzio_client/client"
	"github.com/jonboydell/logzio_client/test_utils"
	"github.com/jonboydell/logzio_client/users"
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

func setupUsersTest() (*users.UsersClient, error, func()) {
	apiToken := "SOME_API_TOKEN"
	underTest, _ := users.New(apiToken)

	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	underTest.BaseUrl = server.URL
	underTest.Client.BaseUrl = server.URL

	return underTest, nil, func() {
		server.Close()
	}
}

func setupUsersIntegrationTest() (*users.UsersClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}

	underTest, err := users.New(apiToken)
	underTest.BaseUrl = client.GetLogzIoBaseUrl()
	underTest.Client.BaseUrl = client.GetLogzIoBaseUrl()
	return underTest, err
}
