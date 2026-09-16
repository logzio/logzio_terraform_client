package tracing_accounts_test

import (
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/logzio/logzio_terraform_client/tracing_accounts"
	"net/http"
	"net/http/httptest"
	"os"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func fixture(path string) string {
	b, err := os.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func setupTracingAccountsTest() (*tracing_accounts.TracingAccountClient, error, func()) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, _ := tracing_accounts.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}

// setupTracingAccountsIntegrationTest authenticates with the consumption account token. Tracing
// accounts live under the consumption account-management API, which rejects a non-consumption owner,
// so they cannot run under the token the non-consumption modules use.
func setupTracingAccountsIntegrationTest() (*tracing_accounts.TracingAccountClient, string, error) {
	apiToken, err := test_utils.GetConsumptionApiToken()
	if err != nil {
		return nil, "", err
	}

	email, err := test_utils.GetLogzioEmail()
	if err != nil {
		return nil, "", err
	}

	underTest, err := tracing_accounts.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, email, err
}

func getCreateOrUpdateTracingAccount(email string) tracing_accounts.CreateOrUpdateTracingAccount {
	return tracing_accounts.CreateOrUpdateTracingAccount{
		AccountName:          "tf_client_test_tracing",
		AuthorizedAccountIds: []int32{},
		MaxDailyGB:           1,
		Email:                email,
	}
}

func validCreateTracingAccount() tracing_accounts.CreateOrUpdateTracingAccount {
	return tracing_accounts.CreateOrUpdateTracingAccount{
		AccountName:          "tf_client_test_tracing",
		AuthorizedAccountIds: []int32{},
		MaxDailyGB:           5,
		Email:                "some@email.test",
	}
}
