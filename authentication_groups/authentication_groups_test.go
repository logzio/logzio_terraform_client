package authentication_groups_test

import (
	"github.com/logzio/logzio_terraform_client/authentication_groups"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

const authGroupsApiBasePath = "/v1/authentication/groups"

func setupAuthenticationGroupsTest() (*authentication_groups.AuthenticationGroupsClient, func(), error) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, _ := authentication_groups.New(apiToken, server.URL)

	return underTest, func() {
		server.Close()
	}, nil
}

func setupAuthenticationGroupsIntegrationTest() (*authentication_groups.AuthenticationGroupsClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}

	underTest, err := authentication_groups.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}