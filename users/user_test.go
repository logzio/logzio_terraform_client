package users_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/logzio/logzio_terraform_client/users"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

const usersApiBasePath = "/v1/user-management"

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

func setupUsersTest() (*users.UsersClient, func(), error) {
	apiToken := "SOME_API_TOKEN"

	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	underTest, err := users.New(apiToken, server.URL)

	return underTest, func() {
		server.Close()
	}, err
}

func setupUsersIntegrationTest() (*users.UsersClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}

	underTest, err := users.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, err
}

func getCreateUser() (users.CreateUpdateUser, error) {
	createUser := users.CreateUpdateUser{
		UserName:  "some_test@test.test",
		FullName:  "user test",
		AccountId: 0,
		Role:      users.UserRoleReadOnly,
	}

	accountId, err := test_utils.GetAccountId()
	if err != nil {
		return createUser, err
	}

	createUser.AccountId = int32(accountId)
	return createUser, nil
}
