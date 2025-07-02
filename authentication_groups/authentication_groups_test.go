package authentication_groups_test

import (
	"github.com/logzio/logzio_terraform_client/authentication_groups"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"net/http"
	"net/http/httptest"
	"os"
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
	b, err := os.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func getCreateGroups() []authentication_groups.AuthenticationGroup {
	return []authentication_groups.AuthenticationGroup{
		{
			Group:    "test_group_admin",
			UserRole: authentication_groups.AuthGroupsUserRoleAdmin,
		},
		{
			Group:    "test_group_readonly",
			UserRole: authentication_groups.AuthGroupsUserRoleReadonly,
		},
		{
			Group:    "test_group_regular",
			UserRole: authentication_groups.AuthGroupsUserRoleRegular,
		},
	}
}

func getEmptyGroup() []authentication_groups.AuthenticationGroup {
	return []authentication_groups.AuthenticationGroup{}
}
