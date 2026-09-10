package tracing_accounts_test

import (
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

func validCreateTracingAccount() tracing_accounts.CreateOrUpdateTracingAccount {
	return tracing_accounts.CreateOrUpdateTracingAccount{
		AccountName:          "tf_client_test_tracing",
		AuthorizedAccountIds: []int32{},
		MaxDailyGB:           5,
		Email:                "some@email.test",
	}
}
